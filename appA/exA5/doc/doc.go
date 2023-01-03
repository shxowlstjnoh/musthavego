package doc

import "fmt"

// CharSize 상수 설명입니다.
const Charsize = 3

const (
	//CharColorRed 빨간색 - iota
	CharColorRed = iota
	//CharColorBlue 파란색
	CharColorBlue
	//CharColorGreen 초록색
	CharColorGreen
)

// PrintDoc 문서 출력 함수 설명입니다.
func PrintDoc() {

}

// TextDoc 문서 작성 시 구조체 예제입니다.
type TextDoc struct {
	//Msg 내부 메시지 입니다.
	Msg string

	//size를 나타냉빈다.
	size int
}

// NewTextDoc 외부로 공개되는 함수 설명입니다.
func NewTextDoc() *TextDoc {
	return &TextDoc{}
}

// PrintDoc 외부로 공개되는 메서드 설명입니다.
func (t *TextDoc) PrintDoc() {
	fmt.Println("This is TextDoc PrintDoc method")
}
