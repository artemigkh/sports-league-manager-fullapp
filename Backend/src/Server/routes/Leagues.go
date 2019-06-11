package routes

import (
	"Server/databaseAccess"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LeaguePermissionChange struct {
	Id            int  `json:"id"`
	Administrator bool `json:"administrator"`
	CreateTeams   bool `json:"createTeams"`
	EditTeams     bool `json:"editTeams"`
	EditGames     bool `json:"editGames"`
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/createLeague
func createNewLeague() gin.HandlerFunc {
	var league databaseAccess.LeagueCore
	return endpoint{
		Entity:        League,
		AccessType:    Create,
		BindData:      func(ctx *gin.Context) bool { return bindAndCheckErr(ctx, &league) },
		IsDataInvalid: func(ctx *gin.Context) (bool, string, error) { return league.ValidateNew() },
		Core: func(ctx *gin.Context) (interface{}, error) {
			leagueId, err := LeaguesDAO.CreateLeague(getUserId(ctx), league)
			return gin.H{"leagueId": leagueId}, err
		},
	}.createEndpointHandler()
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/updateLeague
func updateLeagueInfo_(ctx *gin.Context) {
	hasPermissions, err := Access.League(databaseAccess.Edit, getLeagueId(ctx), getUserId(ctx))
	if accessForbidden(ctx, hasPermissions, err) {
		return
	}

	var league databaseAccess.LeagueCore
	if bindAndCheckErr(ctx, &league) {
		return
	}

	if validator.DataInvalid(ctx, func() (bool, string, error) { return league.ValidateEdit(getLeagueId(ctx)) }) {
		return
	}

	err = LeaguesDAO.UpdateLeague(getLeagueId(ctx), league)
	if checkErr(ctx, err) {
		return
	}

	ctx.Status(http.StatusOK)
}

func updateLeagueInfo() gin.HandlerFunc {
	var league databaseAccess.LeagueCore
	return endpoint{
		Entity:        League,
		AccessType:    Edit,
		BindData:      func(ctx *gin.Context) bool { return bindAndCheckErr(ctx, &league) },
		IsDataInvalid: func(ctx *gin.Context) (bool, string, error) { return league.ValidateEdit(getLeagueId(ctx)) },
		Core: func(ctx *gin.Context) (interface{}, error) {
			return nil, LeaguesDAO.UpdateLeague(getLeagueId(ctx), league)
		},
	}.createEndpointHandler()
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/setLeagueMd
func setLeagueMarkdown() gin.HandlerFunc {
	var markdown databaseAccess.Markdown
	return endpoint{
		Entity:        League,
		AccessType:    Edit,
		BindData:      func(ctx *gin.Context) bool { return bindAndCheckErr(ctx, &markdown) },
		IsDataInvalid: func(ctx *gin.Context) (bool, string, error) { return markdown.Validate() },
		Core: func(ctx *gin.Context) (interface{}, error) {
			oldFile, err := LeaguesDAO.GetMarkdownFile(getLeagueId(ctx))
			if err != nil {
				return nil, err
			}

			fileName, err := MarkdownManager.StoreMarkdown(markdown.Markdown, oldFile)
			if err != nil {
				return nil, err
			}

			return nil, LeaguesDAO.SetMarkdownFile(getLeagueId(ctx), fileName)
		},
	}.createEndpointHandler()
}

func setLeagueMarkdown_(ctx *gin.Context) {
	hasPermissions, err := Access.League(databaseAccess.Edit, getLeagueId(ctx), getUserId(ctx))
	if accessForbidden(ctx, hasPermissions, err) {
		return
	}

	var md databaseAccess.Markdown
	err = ctx.ShouldBindJSON(&md)
	if checkJsonErr(ctx, err) {
		return
	}

	valid, problem, err := md.Validate()
	if dataInvalid(ctx, valid, problem, err) {
		return
	}

	oldFile, err := LeaguesDAO.GetMarkdownFile(getLeagueId(ctx))
	if checkErr(ctx, err) {
		return
	}

	fileName, err := MarkdownManager.StoreMarkdown(md.Markdown, oldFile)
	if checkErr(ctx, err) {
		return
	}

	err = LeaguesDAO.SetMarkdownFile(getLeagueId(ctx), fileName)
	if checkErr(ctx, err) {
		return
	}

	ctx.Status(http.StatusOK)
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/setActiveLeague
func setActiveLeague() gin.HandlerFunc {
	//var // data Type
	return endpoint{
		Entity:     League,
		AccessType: View,
		Core: func(ctx *gin.Context) (interface{}, error) {
			return nil, ElmSessions.SetActiveLeague(ctx, getTargetLeagueId(ctx))
		},
	}.createEndpointHandler()
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/getLeagueInfo
func getActiveLeagueInformation() gin.HandlerFunc {
	return endpoint{
		Entity:     League,
		AccessType: View,
		Core:       func(ctx *gin.Context) (interface{}, error) { return LeaguesDAO.GetLeagueInformation(getLeagueId(ctx)) },
	}.createEndpointHandler()
}
func getActiveLeagueInformation_(ctx *gin.Context) {
	hasPermissions, err := Access.League(databaseAccess.View, getLeagueId(ctx), getUserId(ctx))
	if accessForbidden(ctx, hasPermissions, err) {
		return
	}

	leagueInfo, err := LeaguesDAO.GetLeagueInformation(getLeagueId(ctx))
	if checkErr(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, leagueInfo)
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/joinLeague
func joinActiveLeague(ctx *gin.Context) {
	allowedJoin, err := LeaguesDAO.CanJoinLeague(getLeagueId(ctx), getUserId(ctx))
	if accessForbidden(ctx, allowedJoin, err) {
		return
	}

	err = LeaguesDAO.JoinLeague(getLeagueId(ctx), getUserId(ctx))
	if checkErr(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/getLeagueTeamManagers
func getTeamManagers(ctx *gin.Context) {
	hasPermissions, err := Access.League(databaseAccess.Edit, getLeagueId(ctx), getUserId(ctx))
	if accessForbidden(ctx, hasPermissions, err) {
		return
	}

	teamManagerInfo, err := LeaguesDAO.GetTeamManagerInformation(getLeagueId(ctx))
	if checkErr(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, teamManagerInfo)
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/getPublicLeagues
func getPublicLeagues(ctx *gin.Context) {
	leagueList, err := LeaguesDAO.GetPublicLeagueList()
	if checkErr(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, leagueList)
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/setLeaguePermissions
func setLeaguePermissions() gin.HandlerFunc {
	var permissions databaseAccess.LeaguePermissionsCore
	return endpoint{
		Entity:        League,
		AccessType:    Edit,
		BindData:      func(ctx *gin.Context) bool { return bindAndCheckErr(ctx, &permissions) },
		IsDataInvalid: func(ctx *gin.Context) (bool, string, error) { return permissions.Validate() },
		Core: func(ctx *gin.Context) (interface{}, error) {
			return nil, LeaguesDAO.SetLeaguePermissions(getLeagueId(ctx), getTargetUserId(ctx), permissions)
		},
	}.createEndpointHandler()
}

func setLeaguePermissions_(ctx *gin.Context) {
	hasPermissions, err := Access.League(databaseAccess.Edit, getLeagueId(ctx), getUserId(ctx))
	if accessForbidden(ctx, hasPermissions, err) {
		return
	}

	var permissions databaseAccess.LeaguePermissionsCore
	if bindAndCheckErr(ctx, &permissions) {
		return
	}

	valid, problem, err := permissions.Validate()
	if dataInvalid(ctx, valid, problem, err) {
		return
	}

	err = LeaguesDAO.SetLeaguePermissions(
		getLeagueId(ctx), getTargetUserId(ctx), permissions)
	if checkErr(ctx, err) {
		return
	}

	ctx.Status(http.StatusOK)
}

// https://artemigkh.github.io/ELM-Electronic-League-Manager/#operation/getLeagueMd
func getLeagueMarkdown() gin.HandlerFunc {
	var fileName string
	var err error
	return endpoint{
		Entity:     League,
		AccessType: View,
		BindData: func(ctx *gin.Context) bool {
			fileName, err = LeaguesDAO.GetMarkdownFile(getLeagueId(ctx))
			if checkErr(ctx, err) {
				return true
			} else {
				return false
			}
		},
		Core: func(ctx *gin.Context) (interface{}, error) {
			markdown, err := MarkdownManager.GetMarkdown(fileName)
			return gin.H{"markdown": markdown}, err
		},
	}.createEndpointHandler()
}

func RegisterLeagueHandlers(g *gin.RouterGroup) {
	// League Manage
	g.POST("", createNewLeague())

	g.PUT("", updateLeagueInfo())
	g.PUT("/markdown", setLeagueMarkdown())
	g.GET("/teamManagers", getTeamManagers)                                             //TODO: wrap this one
	g.PUT("/setLeaguePermissions/:userId", storeTargetUserId(), setLeaguePermissions()) //TODO: test this one

	// League Interact
	g.POST("/setActiveLeague/:leagueId", storeTargetLeagueId(), setActiveLeague())
	g.POST("/join", joinActiveLeague)

	// League Information
	g.GET("", getActiveLeagueInformation())
	g.GET("/markdown", getLeagueMarkdown())
	g.GET("/publicLeagues", getPublicLeagues)
}
