package repository

import (
	"github/felipex/kanban-server/internal/database"
	"github/felipex/kanban-server/internal/domain"
)

func NewCategory(c domain.Category) {
	db, err := database.OpenDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO category (name, color) VALUES (?, ?);", c.Name, c.Color)

	if err != nil {
		panic(err)
	}
}

func UpdateCategory(c domain.Category) {
	db, err := database.OpenDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE category SET name = ?, color = ? WHERE id = ?;", c.Name, c.Color, c.ID)

	if err != nil {
		panic(err)
	}
}

func DeleteCategory(id int) error {
	db, err := database.OpenDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM category WHERE id = ?;", id)

	if err != nil {
		panic(err)
	}
	return err
}

func GetCategory(id int) (category domain.Category, err error) {
	db, err := database.OpenDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM category WHERE id = ?;", id)

	if err != nil {
		panic(err)
	}

	for row.Next() {
		err := row.Scan(&category.ID, &category.Name, &category.Color)

		if err != nil {
			panic(err)
		}
	}
	return category, nil

}

func GetAllCategory() ([]domain.Category, error) {
	db, err := database.OpenDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM category;")

	if err != nil {
		panic(err)
	}

	var categories []domain.Category

	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Color)

		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories, nil

}
