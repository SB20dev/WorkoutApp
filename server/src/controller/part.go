package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"workout/src/helper"
	"workout/src/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type PartController struct {
	DB *gorm.DB
}

func getFilterIDs(r *http.Request, paramName string) []int64 {
	q := r.URL.Query()
	if idsStr := q.Get(paramName); len(idsStr) > 0 {
		idStrs := strings.Split(idsStr, "|")
		ids := []int64{}
		for _, idStr := range idStrs {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				continue
			}
			ids = append(ids, id)
		}
		return ids
	}
	return []int64{}
}

//part

func (p *PartController) GetPartsCount(w http.ResponseWriter, r *http.Request, userID int64) error {
	filterClassIDs := getFilterIDs(r, "filter_classes")
	filterStateIDs := getFilterIDs(r, "filter_status")
	count, err := model.FetchPartsCount(p.DB, userID, filterClassIDs, filterStateIDs)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchPartsCountFailure)
	}
	rtn := map[string](interface{}){
		"count": count,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (p *PartController) GetParts(w http.ResponseWriter, r *http.Request, userID int64) error {
	filterClassIDs := getFilterIDs(r, "filter_classes")
	filterStateIDs := getFilterIDs(r, "filter_status")
	db := p.DB.Scopes(helper.Paginate(r))
	parts, err := model.FetchParts(db, userID, filterClassIDs, filterStateIDs)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchPartsFailure)
	}
	rtn := map[string](interface{}){
		"parts": parts,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (p *PartController) PostPart(w http.ResponseWriter, r *http.Request, userID int64) error {
	var body struct {
		ClassID int64  `json:"class_id"`
		Part    string `json:"part"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	} else if body.Part == "" {
		helper.LogError(r, errors.New("Part is empty."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	part := model.Part{
		CommonPart: model.CommonPart{
			ClassID: body.ClassID,
			Part:    body.Part,
		},
		UserID:  userID,
		Deleted: false,
	}

	err = model.CreatePart(p.DB, part)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.CreatePartFailure)
	}
	return nil
}

func (p *PartController) PatchPart(w http.ResponseWriter, r *http.Request, userID int64) error {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		helper.LogError(r, errors.New("id parameter is required."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	var body struct {
		ClassID int64  `json:"class_id"`
		Part    string `json:"part"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	} else if body.Part == "" {
		helper.LogError(r, errors.New("Part is empty."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	err = model.UpdatePart(p.DB, id, body.ClassID, body.Part)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UpdatePartFailure)
	}
	return nil
}

func (p *PartController) DeletePart(w http.ResponseWriter, r *http.Request, userID int64) error {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		helper.LogError(r, errors.New("id parameter is required."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	err = model.DeletePartLogically(p.DB, id)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.DeletePartFailure)
	}
	return nil
}

// class

func (p *PartController) GetClasses(w http.ResponseWriter, r *http.Request, userID int64) error {
	classes, err := model.FetchClasses(p.DB, userID)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchClassesFailure)
	}
	rtn := map[string](interface{}){
		"classes": classes,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (p *PartController) PostClass(w http.ResponseWriter, r *http.Request, userID int64) error {
	var body struct {
		Class string `json:"class"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	} else if body.Class == "" {
		helper.LogError(r, errors.New("Class is empty."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	class := model.Class{
		CommonClass: model.CommonClass{
			Class: body.Class,
		},
		UserID:  userID,
		Deleted: false,
	}

	err = model.CreateClass(p.DB, class)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.CreateClassFailure)
	}
	return nil
}

func (p *PartController) PatchClass(w http.ResponseWriter, r *http.Request, userID int64) error {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		helper.LogError(r, errors.New("id parameter is required."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	var body struct {
		Class string `json:"class"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	} else if body.Class == "" {
		helper.LogError(r, errors.New("Class is empty."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	err = model.UpdateClass(p.DB, id, body.Class)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UpdateClassFailure)
	}
	return nil
}

func (p *PartController) DeleteClass(w http.ResponseWriter, r *http.Request, userID int64) error {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		helper.LogError(r, errors.New("id parameter is required."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	err = model.DeleteClassLogically(p.DB, id)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.DeleteClassFailure)
	}
	return nil
}

// status

func (p *PartController) GetStatus(w http.ResponseWriter, r *http.Request, userID int64) error {
	status, err := model.FetchStatus(p.DB)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchStatusFailure)
	}
	rtn := map[string](interface{}){
		"status": status,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}
