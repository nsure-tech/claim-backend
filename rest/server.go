package rest

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (server *HttpServer) Start() {
	r := SetupRouter()
	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	r.Use(Cors())
	//r.Use(setCROSOptions)

	r.POST("/api/claim", PlaceClaimChain)

	r.GET("/api/arbiter", GetArbiter)

	r.POST("/api/addArbiter", AddArbiter)
	r.POST("/api/pending", PendingArbiter)

	r.GET("/api/claimListNew", GetClaimListNew)
	r.GET("/api/claimListDown", GetClaimListDown)
	r.GET("/api/claimListAdmin", GetClaimListAdmin)

	r.GET("/api/claimApply", GetClaimApply)
	r.POST("/api/claimApply", PlaceClaimApply)

	r.GET("/api/assessment", GetClaimVote)
	r.POST("/api/assessment", PlaceVote)
	r.POST("/api/assessmentAdmin", PlaceVoteAdmin)
	r.GET("/api/assessmentResult", GetClaimResult)

	r.GET("/api/claimList", GetClaimList)
	r.GET("/api/claimListArbiter", GetClaimListByArbiter)

	r.GET("/api/challenge", GetChallengeChaAmount)
	r.POST("/api/challenge", ApplyChallenge)

	r.POST("/api/challengeResult", ChallengeVote)

	r.POST("/api/withdraw", PlaceWithdraw)

	r.GET("/api/asset", AssetHistory)

	r.GET("/api/claimListUser", GetClaimListByUserId)

	// test
	r.POST("/api/withdrawTest", PlaceWithdrawTest)

	r.GET("/api/claimVote", GetClaimVoteOld)
	r.GET("/api/claimResult", GetClaimResult)

	r.GET("/api/claimArbiter", GetClaimByArbiterId)
	r.POST("/api/applyArbiter", ApplyClaimByArbiterId)

	r.GET("/api/login", GetUserRandom)
	r.POST("/api/login", SignIn)

	r.POST("/api/vote", PlaceVoteOld)
	r.POST("/api/challengeResultOld", ChallengeVoteOld)

	private := r.Group("/", checkToken())
	{
		private.GET("/api/loginTest", LoginTest)
	}

	r.POST("/claim/new", PlaceClaimOld)

	r.GET("/api/claim", GetClaimByApply)

	r.POST("/api/deposit", Deposit)
	r.POST("/api/withdraw1", Withdraw)

	r.POST("/api/challengeOld", ApplyChallengeOld)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func setCROSOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	//c.Header("Content-Type", "application/json")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
