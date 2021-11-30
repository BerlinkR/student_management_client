package main

func main(){
	showFunctionList()
		for{
			choose := inputButton()
			switch choose {
			case 1: ShowList()
			case 2: ShowOrder()
			case 3:
				stu := input_student()
				RemoteInsert(stu)
			case 4:
				showFunctionList()
		}
	}

}
