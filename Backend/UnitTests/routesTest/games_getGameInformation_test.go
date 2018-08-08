package routesTest

import (
	"bytes"
	"encoding/json"
	"errors"
	"esports-league-manager/Backend/Server/databaseAccess"
	"esports-league-manager/Backend/Server/routes"
	"esports-league-manager/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"testing"
)

func createGameInfoBody(id int, leagueID, team1ID, team2ID, gameTime,
	winnerID, scoreTeam1, scoreTeam2 int, complete bool) *bytes.Buffer {
	body := databaseAccess.GameInformation{
		Id:         id,
		LeagueID:   leagueID,
		Team1ID:    team1ID,
		Team2ID:    team2ID,
		GameTime:   gameTime,
		Complete:   complete,
		WinnerID:   winnerID,
		ScoreTeam1: scoreTeam1,
		ScoreTeam2: scoreTeam2,
	}
	bodyB, _ := json.Marshal(&body)
	return bytes.NewBuffer(bodyB)
}

func testGetGameInformationSessionError(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(1, errors.New("fake session error"))

	routes.ElmSessions = mockSession

	httpTest(t, nil, "GET", "/1", 500, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testGetGameInformationNoActiveLeague(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(-1, nil)

	routes.ElmSessions = mockSession

	httpTest(t, nil, "GET", "/1", 403, testParams{Error: "noActiveLeague"})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testGetGameInformationNoId(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(1, nil)

	routes.ElmSessions = mockSession

	httpTest(t, nil, "GET", "/", 404, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testGetGameInformationDbError(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(2, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("GetGameInformation", 1, 2).
		Return(nil, errors.New("fake db error"))

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, nil, "GET", "/1", 500, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func testGetGameInformationTeamDoesNotExist(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(2, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("GetGameInformation", 1, 2).
		Return(nil, nil)

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, nil, "GET", "/1", 400, testParams{Error: "gameDoesNotExist"})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func testGetGameInformationNotInt(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(1, nil)

	routes.ElmSessions = mockSession

	httpTest(t, nil, "GET", "/a", 400, testParams{Error: "IdMustBeInteger"})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testGetGameInformationCorrectGetInfo(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(2, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("GetGameInformation", 1, 2).
		Return(&databaseAccess.GameInformation{
			Id:         1,
			LeagueID:   2,
			Team1ID:    4,
			Team2ID:    5,
			GameTime:   1532913359,
			Complete:   true,
			WinnerID:   4,
			ScoreTeam1: 2,
			ScoreTeam2: 1,
		}, nil)

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, nil, "GET", "/1", 200, testParams{ResponseBody: createGameInfoBody(1, 2, 4, 5,
		1532913359, 4, 2, 1, true)})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func Test_GetGameInformation(t *testing.T) {
	//set up router and path to test
	gin.SetMode(gin.ReleaseMode) //opposite of gin.DebugMode to make tests faster by removing logging
	router = gin.New()

	router.Use(routes.Testing_Export_getActiveLeague())
	router.GET("/:id",
		routes.Testing_Export_getUrlId(),
		routes.Testing_Export_getGameInformation)

	t.Run("sessionError", testGetGameInformationSessionError)
	t.Run("noActiveLeague", testGetGameInformationNoActiveLeague)
	t.Run("noId", testGetGameInformationNoId)
	t.Run("IdNotInt", testGetGameInformationNotInt)
	t.Run("teamDoesNotExist", testGetGameInformationTeamDoesNotExist)
	t.Run("dbError", testGetGameInformationDbError)
	t.Run("correctGetInfo", testGetGameInformationCorrectGetInfo)
}
