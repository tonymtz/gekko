package repositories

func openDB(db_options *pg.Options) *pg.DB {
	db := pg.Connect(db_options)

	return db
}

func closeDB(db *pg.DB) {
	db.Close()
}
