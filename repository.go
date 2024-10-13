package winter

import "xorm.io/xorm"

func FindById[T any](engine *xorm.Engine, id int64) (*T, error) {
	entity := new(T)

	if exists, err := engine.ID(id).Get(entity); err != nil || !exists {
		return nil, err
	}

	return entity, nil
}

func FindOne[T any](engine *xorm.Engine, sql string, parameters ...any) (*T, error) {
	entity := new(T)

	if exists, err := engine.SQL(sql, parameters...).Get(entity); err != nil || !exists {
		return nil, err
	}

	return entity, nil
}

func FindAll[T any](engine *xorm.Engine) ([]T, error) {
	entites := make([]T, 0)

	err := engine.Find(&entites)

	return entites, err
}

func FindWithSql[T any](engine *xorm.Engine, sql string, parameters ...any) ([]T, error) {
	entites := make([]T, 0)

	err := engine.SQL(sql, parameters...).Find(&entites)

	return entites, err
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

func Delete[T any](engine *xorm.Engine, bean *T) (int64, error) {
	return engine.Delete(bean)
}

func DeleteWithSession[T any](session *xorm.Session, bean *T) (int64, error) {
	return session.Delete(bean)
}
