package datebase

func CreateDBIfNotExist() error 
{
	res, err := r.DBList().Run(session) 
	if err != nil 	
		{ 
			return err 
		} 

	var dbList []string 
	err = res.All(&dbList) 
	if err != nil 
		{ 
			return err 
		} 

	for _, item := range dbList 	
		{ 
			if item == "kanban" 
			{ 
				return nil 
			} 
		} 

	_, err = r.DBCreate("kanban").Run(session) 
	if err != nil 
	{ 
		return err 
	} 

	return nil 
} 

func CreateTableIfNotExistStage() error 
{ 
	res, err := r.DB("kanban").TableList().Run(session) 
	if err != nil 
	{ 
		return err 
	} 

	var tableList []string 
	err = res.All(&tableList) 
	if err != nil 
		{ 
			return err 
		} 

	for _, item := range tableList 
		{ 
			if item == "Stage" 
				{ 
					return nil 
				} 
		} 

	_, err = r.DB("kanban").TableCreate("Stage", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session) 
	return err 
}  
func CreateTableIfNotExistIssue() error 
{ 
	res, err := r.DB("kanban").TableList().Run(session) 
	if err != nil 
	{ 
		return err 
	} 

	var tableList []string 
	err = res.All(&tableList) 
	if err != nil 
		{ 
			return err 
		} 

	for _, item := range tableList 
		{ 
			if item == "Issue" 
				{ 
					return nil 
				} 
		} 

	_, err = r.DB("kanban").TableCreate("Issue", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session) 
	return err 
}  
