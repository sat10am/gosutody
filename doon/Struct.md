# 구조체
- 필드들의 모음 혹은 묶음을 구조체라고 함
- 구조체에 명명된 구성 요소들은 필드
- 서로 같은 자료형의 자료들을 묶어놓은 것이 배열이라면, 구조체는 서로 다른 자료형의 자료들도 묶을 수 있음
- 구조체는 필드 하나만 가질수도 있다. 그 형태는 아래와 같다
    ```go
    package deadline
    
    import (
        "fmt"
        "time"
    )
  
    type Deadline time.Time

    func (d Deadline) OverDue() bool {
        return time.Time(d).Before(time.Now())
    }

    func ExampleDeadlineOverdue() {
        d1 := Deadline(time.Now().Add(-4 * time.Hour))
        d2 := Deadline(time.Now().Add(4 * time.Hour))
        fmt.Println(d1.OverDue()) // true
        fmt.Println(d2.OverDue()) // false
    }
    
    func main() {
        ExampleDeadlineOverdue()
    }
    ```
- C의 구조체와 대단히 유사하다는 생각이 들었는데, 실제로 호환이 가능했다
    ```go
    package main
    
    /*
    #include <stdlib.h>
    
    typedef struct _PERSON {
    	char *name; // 문자열 포인터
    	int age;    // 정수
    } PERSON;
    
    PERSON* create(char *name, int age) // 메모리 할당 함수 작성
    {
    	PERSON *p = (PERSON *)malloc(sizeof(PERSON)); // PERSON 크기로 메모리 할당
    
    	p->name = name; // 값 설정
    	p->age = age;   // 값 설정
    
    	return p; // 할당한 메모리 리턴
    }
    */
    import "C"
    import (
    	"fmt"
    	"unsafe"
    )
    
    func main() {
    	var p *C.PERSON
    
    	name := C.CString("Maria")
    	age := C.int(20)
    
    	p = C.create(name, age)
    
    	fmt.Println(C.GoString(p.name)) // Maria: char *는 C.GoString 함수로 변환하여 사용
    	fmt.Println(p.age)              // 20
    
    	C.free(unsafe.Pointer(name)) // C.CString으로 만든 문자열은 반드시 해제
    	C.free(unsafe.Pointer(p))    // C 언어의 malloc 함수로 할당한 메모리는 반드시 해제
    }
    ```
- 구조체를 초기화 하는 방법은 여러가지가 있다
    ```go
    package main
    import "fmt"
  
    type Person struct {
        name string
        age int
    }
  
    func newPerson(name string, age int) *Person {
        person := Person{}
        person.name = name
        person.age = age
        return &person
    }
    
    func main() {
        // 1. Struct with empty value
        doon := Person{}
        doon.name = "doon"
        doon.age = 32
        fmt.Println(doon)  
  
        // 2. Struct with initial value
        godori := Person{
            name: "Go dori",
            age: 8,  
        }
        fmt.Println(godori)
  
        // 3. with "new" keyword
        hosung := new(Person)
        hosung.name = "Hosung"
        hosung.age = 7
        
        // 4. with constructor function
        yesdoing := newPerson("Yesdoing", 6)
        fmt.Println(yesdoing)
    }
    ```
- 구조체의 필드는 다른 구조체, 함수 등을 받을 수 있다