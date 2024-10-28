package orm

import "xorm.io/xorm"

func FindById[T any](engine *xorm.Engine, id int64, entity *T) error {
	if exists, err := engine.ID(id).Get(entity); err != nil || !exists {
		return err
	}

	return nil
}

func FindOne[T any](engine *xorm.Engine, entity *T, sql string, parameters ...any) error {
	if exists, err := engine.SQL(sql, parameters...).Get(entity); err != nil || !exists {
		return err
	}

	return nil
}

func FindAll[T any](engine *xorm.Engine, entites *[]T) error {
	return engine.Find(entites)
}

func FindWithSql[T any](engine *xorm.Engine, entites *[]T, sql string, parameters ...any) error {
	return engine.SQL(sql, parameters...).Find(entites)
}

func FindPagination[T any](engine *xorm.Engine, entites *[]T, countSql string, countParameters []any, querySql string, queryParameters []any) (int64, error) {
	total := int64(0)

	if _, err := engine.SQL(countSql, countParameters...).Get(&total); err != nil {
		return total, err
	}

	if total > 0 {
		if err := engine.SQL(querySql, queryParameters...).Find(entites); err != nil {
			return total, err
		}
	}

	return total, nil
}

func Create[T any](engine *xorm.Engine, entity *T) error {
	_, err := engine.Insert(entity)

	return err
}

func CreateWithSession[T any](session *xorm.Session, entity *T) error {
	_, err := session.Insert(entity)

	return err
}

func Update[T any](engine *xorm.Engine, cols []string, entity *T) (int64, error) {
	return engine.Cols(cols...).Update(entity)
}

func UpdateWithSession[T any](session *xorm.Session, cols []string, entity *T) (int64, error) {
	return session.Cols(cols...).Update(entity)
}

func Delete[T any](engine *xorm.Engine, entity *T) (int64, error) {
	return engine.Delete(entity)
}

func DeleteWithSession[T any](session *xorm.Session, entity *T) (int64, error) {
	return session.Delete(entity)
}
