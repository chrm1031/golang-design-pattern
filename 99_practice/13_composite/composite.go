package p_composite

import "fmt"

/*
概念的な例
ファイル・ システムには、 ファイルとフォルダーという 2 種類のオブジェクトが存在します。
ファイルとフォルダーを同じように扱う必要がある場合があります。
ここで、 Composite パターンが役に立ちます。

ファイル・システムで特定のキーワードを使った検索を実行する必要があるとします。
この検索の作業は、 ファイルとフォルダーの両方に適用されます。 ファイルの場合は、 ファイルの内容を調べるだけです。
フォルダーの場合は、 そのフォルダーのすべてのファイルに対してキーワード検索を行います。
*/
/*
コンポーネントのインターフェース
*/
type Component interface {
	search(string)
}

/*
コンテナ・コンポジット
*/
type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

/*
リーフ
*/
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

/*
クライアントコード
*/
func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
