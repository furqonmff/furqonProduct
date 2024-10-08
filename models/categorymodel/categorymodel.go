package categorymodel

import (
	"furqonProduct/Config"
	entities "furqonProduct/Entities"
)

func GetAll() []entities.Category {
	rows, err := Config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreateAt, &category.UpdateAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}
