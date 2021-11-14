package task

import (
	"Jumia_todoList/entity"
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Instance struct {
	Logger *zerolog.Logger
	ORM    *gorm.DB
}

func (l *Instance) Load(logger *zerolog.Logger, orm *gorm.DB) {
	l.ORM = orm
	l.Logger = logger
}

func (l *Instance) Create(i *entity.Task, str []string) error {
	tags := make([]entity.Tag, 0)

	for _, v := range str {
		tags = append(tags, entity.Tag{Name: v})
	}
	i.Tags = tags
	return l.ORM.Create(i).Error
}

func (l *Instance) Update(i *entity.Task, id uint) error {
	return l.ORM.Model(&entity.Task{}).Where("id=?", id).Updates(i).Error
}

func (l *Instance) Delete(id int) error {
	return l.ORM.Unscoped().Delete(&entity.Task{}, id).Error
}

func (l *Instance) FilterList(listId int, tags []string) ([]entity.Task, error) {
	query := "select tasks.id,tasks.title, tasks.description,tasks.due, tasks.completed, array_agg(tags.id) As tag_ids ,array_agg(tags.name) As tag_names  from (select * from tasks where list_id =@list_id )tasks left join tags on tasks.id = tags.task_id where tags.name in @tags group by tasks.id,tasks.title, tasks.description,tasks.due, tasks.completed"
	result, err := l.ORM.Raw(query, map[string]interface{}{"list_id": listId, "tags": tags}).Rows()

	if err != nil {
		fmt.Println("errrrr", err)
		return nil, err
	}
	defer result.Close()

	tasks := make([]entity.Task, 0)

	for result.Next() {
		var task entity.Task
		var gpTags groupedTags
		fmt.Println(result.Columns())
		err := result.Scan(&task.ID,
			&task.Title,
			&task.Description,
			&task.Due,
			&task.Completed,
			&gpTags.TagIds,
			&gpTags.TagNames,
		)
		if err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		tags := mapTags(gpTags)
		task.Tags = tags
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (l *Instance) FilterInAllList(tags []string) ([]entity.Task, error) {
	query := "select tasks.id,tasks.title, tasks.description,tasks.due, tasks.completed, array_agg(tags.id) As tag_ids ,array_agg(tags.name) As tag_names  from tasks left join tags on tasks.id = tags.task_id where tags.name in @tags group by tasks.id,tasks.title, tasks.description,tasks.due, tasks.completed"
	result, err := l.ORM.Raw(query, map[string]interface{}{"tags": tags}).Rows()

	if err != nil {
		return nil, err
	}
	defer result.Close()

	tasks := make([]entity.Task, 0)

	for result.Next() {
		var task entity.Task
		var gpTags groupedTags
		fmt.Println(result.Columns())
		err := result.Scan(&task.ID,
			&task.Title,
			&task.Description,
			&task.Due,
			&task.Completed,
			&gpTags.TagIds,
			&gpTags.TagNames,
		)
		if err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		tags := mapTags(gpTags)
		task.Tags = tags
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (l *Instance) Get(listId int) ([]entity.Task, error) {
	query := "select tasks.id,tasks.title, tasks.description,tasks.due, tasks.completed, array_agg(tags.id) As tag_ids ,array_agg(tags.name) As tag_names  from (select * from tasks where list_id =@id )tasks left join tags on tasks.id = tags.task_id  group by tasks.id,tasks.title, tasks.description,tasks.due, tasks.completed"
	result, err := l.ORM.Raw(query, map[string]interface{}{"id": listId}).Rows()

	if err != nil {
		return nil, err
	}
	defer result.Close()

	tasks := make([]entity.Task, 0)

	for result.Next() {
		var task entity.Task
		var gpTags groupedTags
		fmt.Println(result.Columns())
		err := result.Scan(&task.ID,
			&task.Title,
			&task.Description,
			&task.Due,
			&task.Completed,
			&gpTags.TagIds,
			&gpTags.TagNames,
		)
		if err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		tags := mapTags(gpTags)
		task.Tags = tags
		tasks = append(tasks, task)
	}
	return tasks, nil
}
