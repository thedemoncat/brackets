package raw_parser

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestBucketsParser_nextNodeText(t *testing.T) {

	tests := []struct {
		name string
		text string
		want string
	}{
		{
			"simple",
			"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},",
			"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				rd: bufio.NewReader(strings.NewReader(tt.text)),
			}
			if got := p.nextNodeText(); string(got) != tt.want {
				t.Errorf("nextNodeText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucketsParser_NextNode(t *testing.T) {
	tests := []struct {
		name string
		text string
		want *bracketsNode
	}{
		{
			"simple",
			"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},",
			&bracketsNode{
				Nodes: BracketNodes{
					newValueNode("20200412134348"),
					newValueNode("N"),
					bracketsNode{
						Nodes: BracketNodes{
							newValueNode("0"),
							newValueNode("0"),
						},
					},
					newValueNode("1"),
					newValueNode("1"),
					newValueNode("1"),
					newValueNode("1"),
					newValueNode("1"),
					newValueNode("1"),
					newValueNode(""),
					newValueNode("0"),
					bracketsNode{
						Nodes: BracketNodes{
							newValueNode("U"),
						},
					},
					newValueNode(""),
					newValueNode("1"),
					newValueNode("1"),
					newValueNode("0"),
					newValueNode("1"),
					newValueNode("0"),
					bracketsNode{
						Nodes: BracketNodes{
							newValueNode("0"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				rd: bufio.NewReader(strings.NewReader(tt.text)),
			}
			if got := p.NextNode(); strings.EqualFold(got.String(), tt.want.String()) {
				t.Errorf("NextNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucketsParser_ReadAllNodes(t *testing.T) {

	tests := []struct {
		name string
		text string
		want int
	}{
		{
			"simple",
			"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},",
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				rd: bufio.NewReader(strings.NewReader(tt.text)),
			}
			if got := p.ReadAllNodes(); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("ReadAllNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
