package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
	"workout/src/helper"
	"workout/src/model"

	"gorm.io/gorm"
)

type CommitmentController struct {
	DB *gorm.DB
}

func (c *CommitmentController) GetTotalScore(w http.ResponseWriter, r *http.Request, userID string) error {
	totalScore, err := model.FetchTotalCommitmentScore(c.DB, userID)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchTotalCommitmentScoreFailure)
	}

	rtn := map[string]int{
		"total": totalScore,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *CommitmentController) GetCount(w http.ResponseWriter, r *http.Request, userID string) error {
	count, err := model.FetchCommitmentCount(c.DB, userID)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchCommitmentCountFailure)
	}

	rtn := map[string]int64{
		"count": count,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *CommitmentController) GetHistory(w http.ResponseWriter, r *http.Request, userID string) error {
	// offset, num
	q := r.URL.Query()
	offset, offsetErr := strconv.Atoi(q.Get("offset"))
	if offsetErr != nil {
		helper.LogError(r, offsetErr, nil)
	}
	num, numErr := strconv.Atoi(q.Get("num"))
	if numErr != nil {
		helper.LogError(r, numErr, nil)
	}
	if offsetErr != nil || numErr != nil {
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidQueryParameter)
	}

	histories, err := model.FetchCommitmentHistories(c.DB, userID, offset, num)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchCommitmentHistoryFailure)
	}

	rtn := map[string]([]model.Commitment){
		"histories": histories,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *CommitmentController) GetDetail(w http.ResponseWriter, r *http.Request, userID string) error {
	q := r.URL.Query()
	commitment_id, err := strconv.Atoi(q.Get("commitment_id"))
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidQueryParameter)
	}

	commitmentDetail, err := model.FetchCommitmentDetail(c.DB, commitment_id)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchCommitmentDetailFailure)
	}

	return helper.JSON(w, http.StatusOK, commitmentDetail)
}

func (c *CommitmentController) Post(w http.ResponseWriter, r *http.Request, userID string) error {
	var body struct {
		menus []model.CommitmentMenu `json:"menus"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	} else if body.menus == nil || len(body.menus) == 0 {
		helper.LogError(r, errors.New("Menu is nil or its length equals zero."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	// score
	score := calcCommitmentScore(body.menus)

	// commitment
	commitment := model.Commitment{
		UserID:    userID,
		Score:     score,
		Committed: time.Now(),
	}

	err = model.CreateCommitment(c.DB, &commitment, body.menus)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.CreationCommitmentFailure)
	}
	return helper.JSON(w, http.StatusOK, nil)
}

func calcCommitmentScore(menus []model.CommitmentMenu) int {
	// TODO：計算式については要検討
	return len(menus)
}
