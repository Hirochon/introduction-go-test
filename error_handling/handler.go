package error_handling

import (
	"fmt"
	"log"
)

func downProcess(isPanic bool) {
	if isPanic {
		panic("illegal processing")
	}
	log.Fatal("illegal processing")
}

func initializeSQLServer(isDown bool) {
	if isDown {
		panic("はにゃ？")
	}
}

func isAgeSmallerThanOrEqualToZero(age int) bool {
	return age < 0
}

func isAgeGreaterThanGuinnessWorldRecords(age int) bool {
	return age > 150
}

func RequestValidation(age int) error {
	if isAgeSmallerThanOrEqualToZero(age) {
		return fmt.Errorf("age is smaller than or equal to zero")
	}
	if isAgeGreaterThanGuinnessWorldRecords(age) {
		return fmt.Errorf("age is greater than Guinness World Records")
	}
	return nil
}

//func RequestValidationPatternTryCatch(age int) {
//	try{
//		isAgeSmallerThanOrEqualToZero(age)
//		isAgeGreaterThanGuinnessWorldRecords(age)
//	}
//	catch(e) {
//		case(AgeSmallerThanOrEqualToZeroException) => exception("age is smaller than or equal to zero")
//		case(AgeGreaterThanGuinnessWorldRecordsException) => exception("age is greater than Guinness World Records")
//	}
//	finaly {
//		close()
//	}
//}
