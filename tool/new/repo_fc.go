package new

func init() {
	Files = append(Files, repofc1, repofc2)
}

var (
	repofc1 = &FileContent{
		FileName: "user_repo.go",
		Dir:      "internal/infra/repo",
		Content: `package repo

import (
	"context"
	"github.com/jukylin/esim/log"
	"{{.ProPath}}{{.ServerName}}/internal/domain/user/entity"
	"{{.ProPath}}{{.ServerName}}/internal/infra/dao"
)

type UserRepo interface {
	FindByUserName(context.Context, string) entity.User
}

type userRepo struct {
	log log.Logger

	userDao *dao.UserDao
}

func NewDBUserRepo(logger log.Logger) UserRepo {
	repo := &userRepo{
		log: logger,
	}

	if repo.userDao == nil {
		repo.userDao = dao.NewUserDao()
	}

	return repo
}

func (this *userRepo) FindByUserName(ctx context.Context, username string) entity.User {
	var user entity.User
	var err error

	user, err = this.userDao.Find(ctx, "*", "username = ? ", username)

	if err != nil {
		this.log.Errorf(err.Error())
		return user
	}

	return user
}
`,
	}

	repofc2 = &FileContent{
		FileName: "integration_test.go",
		Dir:      "internal/infra/repo",
		Content: `package repo

import (
	"os"
	"testing"
	"github.com/jukylin/esim/mysql"
	"github.com/jukylin/esim/config"
)

var mysqlClient *mysql.Client

func TestMain(m *testing.M) {
	clientOptions := mysql.ClientOptions{}

	options := config.ViperConfOptions{}
	confFile := "../../../conf/dev.yaml"
	file := []string{confFile}
	conf := config.NewViperConfig(options.WithConfigType("yaml"),
		options.WithConfFile(file))

	mysqlClient = mysql.NewClient(clientOptions.WithConf(conf))

	setUp()

	code := m.Run()

	tearDown()

	os.Exit(code)
}

func setUp()  {

}

func tearDown()  {
	mysqlClient.Close()
}
`,
	}
)
