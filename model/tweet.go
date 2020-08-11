package model

import "fmt"

type Tweet struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

func (t *Tweet) GetItem(id string) error {
	fmt.Println("test")
	*t = Tweet{User: id, Content: "Hello world"}
	return nil
}

/*func (t *Tweet) GetList() error {
    *t = []Tweet{
        { User: "mishiro", Content: "Hello world!" },
        { User: "reiji", Content: "See you tomorrow!" },
    }
    return nil
}
*/
