package integration

import (
	"belajar-echo/database"
	"belajar-echo/model"
	"belajar-echo/repository"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite

	SqlDb *sql.DB
	Repo  repository.ProductRepoCont

	Tx *sql.Tx
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &RepositoryTestSuite{})
}

func (s *RepositoryTestSuite) SetupSuite() {
	db, err := database.SetupSQLDatabase()
	s.Require().NoError(err)
	s.SqlDb = db

	s.Require().NoError(err)

	s.Repo = repository.NewProductRepo(s.SqlDb)
}

func (s *RepositoryTestSuite) TearDownSuite() {
	s.SqlDb.Close()
}
func (s *RepositoryTestSuite) SetupTest() {
	var err error
	s.Tx, err = s.SqlDb.Begin()

	if err != nil {
		s.Error(err)
	}
}

func (s *RepositoryTestSuite) TearDownTest() {
	s.Tx.Rollback()
}

func (s *RepositoryTestSuite) TestGetAll() {
	result, err := s.Repo.GetAll()
	s.Assert().NoError(err)
	s.Assert().NotEmpty(result)
}

func (s *RepositoryTestSuite) TestInsert() {
	produk := model.Product{
		Name:  "Produk Test",
		Price: 50000,
	}

	result, err := s.Repo.Insert(produk)
	s.Require().NoError(err)
	s.Assertions.Equal(produk, result, "Tidak sesuai")

}
