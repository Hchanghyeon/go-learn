package main

import "fmt"

func Datatype(){
	// bool
	var isGame bool = true
	
	// string
	var interLiteral string = "changhyeon"
	rawLiteral := `hello\n 
					hello\n
					hello` // 별도로 해석되지 않괴 Raw String 그대로의 값 유지

	name2 := "yiseul" // 문자열 안에 \n이 있을 경우 NewLine으로 해석


	// integer 정수
	var type1 int = 2147483647
	var type2 int8 = 127
	var type3 int16 = 32767
	var type4 int32 = 2147483647
	var type5 int64 = 9223372036854775807

	var type6 uint = 4294967295
	var type7 uint8 = 255
	var type8 uint16 = 65535
	var type9 uint32 = 4294967295
	var type10 uint64 = 18446744073709551615
	var type11 uintptr // 메모리 주소(포인터)의 정수형 

	// decimal 소수
    const decimal1 float32 = 3.14159        // 32비트 부동소수점
    const decimal2 float64 = 2.718281828459 // 64비트 부동소수점 (더 정밀)
    const decimal3 complex64 = 1 + 2i       // 32비트 실수와 32비트 허수로 구성된 복소수
    const decimal4 complex128 = 3 + 4i      // 64비트 실수와 64비트 허수로 구성된 복소수
	

	// 기타 타입
	const byte1 byte = 255 // 'A'
	const rune1 rune = '가' // 유니코드 값 44032가 저장 됨

	// 형 변환
	
	var type1Decimal float32 = float32(type1)
	var type2uint uint = uint(type2)

	bytes := []byte(name2)


	fmt.Println("------------datatype------------")
}
