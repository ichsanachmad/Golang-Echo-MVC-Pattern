package model

import (
	"Golang-Echo-MVC-Pattern/entity"
	"database/sql"
)

type ExampleModel struct{}

// Get Example Post
func (ExampleModel ExampleModel) GetPosts(secret string, db *sql.DB) []entity.ExampleEntity {
	posts := []entity.ExampleEntity{
		{
			Title:       "NewsOne",
			Description: "NewsOneDescription",
		},
		{
			Title:       "NewsTwo",
			Description: "NewsTwoDescription",
		},
		{
			Title:       secret,
			Description: "NewsTwoDescription",
		},
	}

	return posts
}

// Example For Insert Update and Delete
//func (ExampleModel Model) InsertDb(exam ExampleStruct, db *sql.DB) bool {
//
//	sqlStatement := "INSERT INTO daerah_kecamatan (kecamatan) " +
//		"VALUES ($1)"
//	res, err := db.Exec(sqlStatement,
//		exam.Kecamatan,
//	)
//
//	if err != nil {
//		fmt.Println(err)
//		return false
//	} else {
//		return true
//	}
//
//}

// Example For Get From DB
//func (ExampleModel Model) GetDb(db *sql.DB) Exams {
//
//	sqlStatement := "SELECT * FROM daerah_kecamatan " +
//					"ORDER BY daerah_kecamatan.id ASC"
//	res, err := db.Query(sqlStatement)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	result := Exams{}
//	totalExam := ExamDataTotal{}
//
//	for res.Next() {
//		exam := ExamData{}
//		err2 := res.Scan(&exam.Id, &exam.Kecamatan)
//		// Exit if we get an error
//		if err2 != nil {
//			fmt.Println(err2)
//		}
//		result.Exams = append(result.Exams, exam)
//	}
//	defer res.Close()
//
//	result.Total = totalExam
//
//	return result
//}
