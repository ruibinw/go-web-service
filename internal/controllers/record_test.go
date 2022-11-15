package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"git.epam.com/ryan_wang/go-web-service/internal/dto"
	customErrors "git.epam.com/ryan_wang/go-web-service/internal/errors"
	"git.epam.com/ryan_wang/go-web-service/internal/models"
	"git.epam.com/ryan_wang/go-web-service/internal/services/mocks"
	"git.epam.com/ryan_wang/go-web-service/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type testCase struct {
	name                string
	httpMethod          string
	requestBody         any
	pathParams          string
	queryParams         string
	expectedServiceCall bool
	expectedSuccess     bool
	expectedStatusCode  int
	expectedData        any
	expectedError       error
}

func TestRecordController_Create(t *testing.T) {
	testCases := []testCase{
		{
			name:                "Create success",
			httpMethod:          http.MethodPost,
			requestBody:         newCreateReqDto("/url1", "name1", "desc1"),
			expectedServiceCall: true,
			expectedSuccess:     true,
			expectedStatusCode:  http.StatusCreated,
			expectedData:        newRecord(1, "/url1", "name1", "desc1"),
			expectedError:       nil,
		},
		{
			name:                "Create error - url is required",
			httpMethod:          http.MethodPost,
			requestBody:         newCreateReqDto("", "name1", ""),
			expectedServiceCall: false,
			expectedSuccess:     false,
			expectedStatusCode:  http.StatusBadRequest,
			expectedData:        nil,
			expectedError:       errors.New("Url: input is required"),
		},
		{
			name:                "Create error - url is required",
			httpMethod:          http.MethodPost,
			requestBody:         newCreateReqDto("/url1", "", ""),
			expectedServiceCall: false,
			expectedSuccess:     false,
			expectedStatusCode:  http.StatusBadRequest,
			expectedData:        nil,
			expectedError:       errors.New("DisplayName: input is required"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqContext, responseRec, expectedBody := setUpEchoHttpTest(tc)

			// setup Mock Service
			mockService := mocks.NewMockRecordService(gomock.NewController(t))
			if tc.expectedServiceCall {
				mockService.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(tc.expectedData, tc.expectedError)
			}

			// test target controller method and assert
			testController := NewRecordController(mockService)
			if assert.NoError(t, testController.Create(reqContext)) {
				assert.Equal(t, tc.expectedStatusCode, responseRec.Code)
				assert.Equal(t, expectedBody, responseRec.Body.String())
			}
		})
	}

}

func TestRecordController_Update(t *testing.T) {
	testCases := []testCase{
		{
			name:                "Update success",
			httpMethod:          http.MethodPut,
			requestBody:         newUpdateReqDto("/url1", "name1", "desc1"),
			pathParams:          "id:1",
			expectedServiceCall: true,
			expectedSuccess:     true,
			expectedStatusCode:  http.StatusOK,
			expectedData:        newRecord(1, "/url1", "name1", "desc1"),
		},
		{
			name:                "Update error - record not found",
			httpMethod:          http.MethodPut,
			requestBody:         newUpdateReqDto("/url1", "name1", "desc1"),
			pathParams:          "id:1",
			expectedServiceCall: true,
			expectedSuccess:     false,
			expectedStatusCode:  http.StatusNotFound,
			expectedError:       customErrors.NewRecordNotFoundError(1),
		},
		{
			name:                "Update error - display is required",
			httpMethod:          http.MethodPut,
			requestBody:         newUpdateReqDto("/url1", "", ""),
			pathParams:          "id:1",
			expectedServiceCall: false,
			expectedSuccess:     false,
			expectedStatusCode:  http.StatusBadRequest,
			expectedError:       errors.New("DisplayName: input is required"),
		},
		{
			name:                "Update error - url is required",
			httpMethod:          http.MethodPut,
			requestBody:         newUpdateReqDto("", "name1", ""),
			pathParams:          "id:1",
			expectedServiceCall: false,
			expectedSuccess:     false,
			expectedStatusCode:  http.StatusBadRequest,
			expectedError:       errors.New("Url: input is required"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqContext, responseRec, expectedBody := setUpEchoHttpTest(tc)

			// setup Mock Service
			mockService := mocks.NewMockRecordService(gomock.NewController(t))
			if tc.expectedServiceCall {
				mockService.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(tc.expectedData, tc.expectedError)
			}

			// test target controller method and assert
			testController := NewRecordController(mockService)
			if assert.NoError(t, testController.Update(reqContext)) {
				assert.Equal(t, tc.expectedStatusCode, responseRec.Code)
				assert.Equal(t, expectedBody, responseRec.Body.String())
			}
		})
	}
}

func TestRecordController_Delete(t *testing.T) {
	testCases := []testCase{
		{
			name:                "Delete success",
			httpMethod:          http.MethodDelete,
			pathParams:          "id:1",
			expectedServiceCall: true,
			expectedSuccess:     true,
			expectedStatusCode:  http.StatusNoContent,
			expectedError:       nil,
		},
		{
			name:                "Delete error - record not found",
			httpMethod:          http.MethodDelete,
			pathParams:          "id:1",
			expectedServiceCall: true,
			expectedSuccess:     false,
			expectedStatusCode:  http.StatusNotFound,
			expectedError:       customErrors.NewRecordNotFoundError(1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqContext, responseRec, expectedBody := setUpEchoHttpTest(tc)

			// setup Mock Service
			mockService := mocks.NewMockRecordService(gomock.NewController(t))
			if tc.expectedServiceCall {
				mockService.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(tc.expectedError)
			}

			// test target controller method and assert
			testController := NewRecordController(mockService)
			if assert.NoError(t, testController.Delete(reqContext)) {
				assert.Equal(t, tc.expectedStatusCode, responseRec.Code)
				assert.Equal(t, expectedBody, responseRec.Body.String())
			}
		})
	}
}

func TestRecordController_Get(t *testing.T) {
	testCases := []testCase{
		{
			name:                "Get success",
			httpMethod:          http.MethodGet,
			pathParams:          "id:1",
			expectedServiceCall: true,
			expectedSuccess:     true,
			expectedStatusCode:  http.StatusOK,
			expectedData:        newRecord(1, "/url1", "name1", "desc1"),
		},
		{
			name:                "Get error - record not found",
			httpMethod:          http.MethodGet,
			pathParams:          "id:1",
			expectedServiceCall: true,
			expectedSuccess:     false,
			expectedStatusCode:  http.StatusNotFound,
			expectedError:       customErrors.NewRecordNotFoundError(1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqContext, responseRec, expectedBody := setUpEchoHttpTest(tc)

			// setup Mock Service
			mockService := mocks.NewMockRecordService(gomock.NewController(t))
			if tc.expectedServiceCall {
				mockService.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(tc.expectedData, tc.expectedError)
			}

			// test target controller method and assert
			testController := NewRecordController(mockService)
			if assert.NoError(t, testController.Get(reqContext)) {
				assert.Equal(t, tc.expectedStatusCode, responseRec.Code)
				assert.Equal(t, expectedBody, responseRec.Body.String())
			}
		})
	}
}

func TestRecordController_Query(t *testing.T) {
	testCases := []testCase{
		{
			name:                "Query success - query displayName",
			httpMethod:          http.MethodGet,
			queryParams:         "display_name=go",
			expectedServiceCall: true,
			expectedSuccess:     true,
			expectedStatusCode:  http.StatusOK,
			expectedData: []*models.Record{
				newRecord(1, "/url1", "go language", "desc1"),
				newRecord(2, "/url2", "learning go", "desc2"),
			},
		},
		{
			name:                "Query success - query page",
			httpMethod:          http.MethodGet,
			queryParams:         "PageNum=0&PageSize=5",
			expectedServiceCall: true,
			expectedSuccess:     true,
			expectedStatusCode:  http.StatusOK,
			expectedData: []*models.Record{
				newRecord(1, "/url1", "name1", "desc1"),
				newRecord(2, "/url2", "name2", "desc2"),
				newRecord(3, "/url3", "name3", "desc3"),
				newRecord(4, "/url4", "name4", "desc4"),
				newRecord(5, "/url5", "name5", "desc5"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqContext, responseRec, expectedBody := setUpEchoHttpTest(tc)

			// setup Mock Service
			mockService := mocks.NewMockRecordService(gomock.NewController(t))
			if tc.expectedServiceCall {
				mockService.EXPECT().
					Query(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(tc.expectedData, tc.expectedError)
			}

			// test target controller method and assert
			testController := NewRecordController(mockService)
			if assert.NoError(t, testController.Query(reqContext)) {
				assert.Equal(t, tc.expectedStatusCode, responseRec.Code)
				assert.Equal(t, expectedBody, responseRec.Body.String())
			}
		})
	}
}

func newRecord(id int64, url string, name string, desc string) *models.Record {
	return &models.Record{
		ID:          id,
		Url:         url,
		DisplayName: name,
		Description: desc,
		CreatedTime: time.Time{},
		UpdatedTime: time.Time{},
	}
}

func newCreateReqDto(url string, name string, desc string) *dto.CreateRecordRequest {
	return &dto.CreateRecordRequest{
		Url:         url,
		DisplayName: name,
		Description: desc,
	}
}

func newUpdateReqDto(url string, name string, desc string) *dto.UpdateRecordRequest {
	return &dto.UpdateRecordRequest{
		Url:         url,
		DisplayName: name,
		Description: desc,
	}
}

func setUpEchoHttpTest(tc testCase) (echo.Context, *httptest.ResponseRecorder, string) {
	// setup http test context
	e := echo.New()
	e.Validator = utils.NewRequestValidator(validator.New())
	var queryParams string
	if len(tc.queryParams) > 0 {
		queryParams = "?" + tc.queryParams
	}
	body, _ := json.Marshal(tc.requestBody)
	reqBody := bytes.NewReader(body)
	req := httptest.NewRequest(tc.httpMethod, "/records"+queryParams, reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	if len(tc.pathParams) > 0 {
		param := strings.Split(tc.pathParams, ":")
		ctx.SetParamNames(param[0])
		ctx.SetParamValues(param[1])
	}

	// expected response body
	var responseBody utils.ResponseBody
	if tc.expectedSuccess {
		responseBody = utils.ResponseBody{
			Success: true,
			Data:    tc.expectedData,
		}
	} else {
		responseBody = utils.ResponseBody{
			Success: false,
			Errors:  tc.expectedError.Error(),
		}
	}
	body, _ = json.Marshal(responseBody)
	//adds a new line at the end of response to match with echo context encoder result
	expectedBody := string(body) + "\n"

	return ctx, rec, expectedBody
}
