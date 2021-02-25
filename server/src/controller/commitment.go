package controller

import (
	"WorkoutApp/server/src/helper"
	"WorkoutApp/server/src/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type CommitmentController struct {
	DB *gorm.DB
}

func (c *CommitmentController) GetTotalScore(w http.ResponseWriter, r *http.Request) error {
	userID, ok := helper.GetClaim(r, "id").(string)
	if !ok {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to read user id from context")
	}
	totalScore := model.FetchTotalCommitmentScore(c.DB, userID)

	rtn := map[string]int{
		"total": totalScore,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *CommitmentController) GetCount(w http.ResponseWriter, r *http.Request) error {
	userID, ok := helper.GetClaim(r, "id").(string)
	if !ok {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to read user id from context")
	}
	count := model.FetchCommitmentCount(c.DB, userID)

	rtn := map[string]int{
		"count": count,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *CommitmentController) GetHistory(w http.ResponseWriter, r *http.Request) error {
	// userID
	userID, ok := helper.GetClaim(r, "id").(string)
	if !ok {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to read user id from context.")
	}
	// offset, num
	q := r.URL.Query()
	offset, offsetErr := strconv.Atoi(q.Get("offset"))
	num, numErr := strconv.Atoi(q.Get("num"))
	if offsetErr != nil || numErr != nil {
		return helper.CreateHTTPError(http.StatusBadRequest, "query parameter is invalid. convertion failed.")
	}

	histories, err := model.FetchCommitmentHistories(c.DB, userID, offset, num)
	if err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to fetch histories.")
	}

	rtn := map[string]([]model.Commitment){
		"histories": histories,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *CommitmentController) GetDetail(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query()
	commitment_id, err := strconv.Atoi(q.Get("commitment_id"))
	if err != nil {
		return helper.CreateHTTPError(http.StatusBadRequest, "query parameter is invalid. convertion failed.")
	}

	commitmentDetail, err := model.FetchCommitmentDetail(c.DB, commitment_id)
	if err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to fetch commitment details")
	}

	return helper.JSON(w, http.StatusOK, commitmentDetail)
}

func (c *CommitmentController) Post(w http.ResponseWriter, r *http.Request) error {
	// userID
	userID, ok := helper.GetClaim(r, "id").(string)
	if !ok {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to read user id from context.")
	}

	var body struct {
		menus []model.CommitmentMenu `json:menus`
	}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		return helper.CreateHTTPError(http.StatusBadRequest, "request body is invalid. parse failed")
	} else if body.menus != nil && len(body.menus) == 0 {
		return helper.CreateHTTPError(http.StatusBadRequest, "menu length equals zero or nil.")
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
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to create commitment.")
	}

	return nil
}

func calcCommitmentScore(menus []model.CommitmentMenu) int {
	// TODO：計算式については要検討
	return len(menus)
}
