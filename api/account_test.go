package api

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/golang/mock/gomock"
// 	mockdb "github.com/manindhra1412/simple_bank/db/mock"
// 	db "github.com/manindhra1412/simple_bank/db/sqlc"
// 	"github.com/manindhra1412/simple_bank/token"
// 	"github.com/manindhra1412/simple_bank/util"
// 	"github.com/stretchr/testify/require"
// )

// func TestGetAccountAPI(t *testing.T) {
// 	user, _ := randomUser(t)
// 	account := randomAccount(user.Username)

// 	testCases := []struct {
// 		name          string
// 		accountID     int64
// 		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:      "OK",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(account, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchAccount(t, recorder.Body, account)
// 			},
// 		},
// 		{
// 			name:      "UnauthorizedUser",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user", time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(account, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "NoAuthorization",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "NotFound",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
// 			},

// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(db.Account{}, db.ErrRecordNotFound)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "InternalError",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(db.Account{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "InvalidID",
// 			accountID: 0,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			url := fmt.Sprintf("/accounts/%d", tc.accountID)
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			tc.setupAuth(t, request, server.tokenMaker)
// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func randomUser(t *testing.T) (user db.User, password string) {
// 	password = util.RandomString(6)
// 	hashedPassword, err := util.HashedPassword(password)
// 	require.NoError(t, err)

// 	user = db.User{
// 		Username:       util.RandomOwner(),
// 		HashedPassword: hashedPassword,
// 		FullName:       util.RandomOwner(),
// 		Email:          util.RandomEmail(),
// 	}
// 	return
// }

// func randomAccount(owner string) db.Account {
// 	return db.Account{
// 		ID:       util.RandomInt(1, 1000),
// 		Owner:    owner,
// 		Balance:  util.RandomMoney(),
// 		Currency: util.RandomCurrency(),
// 	}
// }
