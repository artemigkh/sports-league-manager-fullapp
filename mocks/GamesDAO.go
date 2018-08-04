// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import databaseAccess "esports-league-manager/Backend/Server/databaseAccess"
import mock "github.com/stretchr/testify/mock"

// GamesDAO is an autogenerated mock type for the GamesDAO type
type GamesDAO struct {
	mock.Mock
}

// CreateGame provides a mock function with given fields: leagueID, team1ID, team2ID, gameTime
func (_m *GamesDAO) CreateGame(leagueID int, team1ID int, team2ID int, gameTime int) (int, error) {
	ret := _m.Called(leagueID, team1ID, team2ID, gameTime)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int, int, int) int); ok {
		r0 = rf(leagueID, team1ID, team2ID, gameTime)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int, int) error); ok {
		r1 = rf(leagueID, team1ID, team2ID, gameTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoesExistConflict provides a mock function with given fields: team1ID, team2ID, gameTime
func (_m *GamesDAO) DoesExistConflict(team1ID int, team2ID int, gameTime int) (bool, error) {
	ret := _m.Called(team1ID, team2ID, gameTime)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int, int) bool); ok {
		r0 = rf(team1ID, team2ID, gameTime)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(team1ID, team2ID, gameTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGameInformation provides a mock function with given fields: gameID, leagueID
func (_m *GamesDAO) GetGameInformation(gameID int, leagueID int) (*databaseAccess.GameInformation, error) {
	ret := _m.Called(gameID, leagueID)

	var r0 *databaseAccess.GameInformation
	if rf, ok := ret.Get(0).(func(int, int) *databaseAccess.GameInformation); ok {
		r0 = rf(gameID, leagueID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*databaseAccess.GameInformation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(gameID, leagueID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}