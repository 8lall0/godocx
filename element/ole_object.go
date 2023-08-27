package element

type OleObject struct {
	Source          string
	Style           string
	Icon            string
	ImageRelationId int
	MediaRelation   bool
}

func pippo() {
	//$supportedTypes = ['xls', 'doc', 'ppt', 'xlsx', 'docx', 'pptx'];
	//$pathInfo = pathinfo($source);
	//
	//if (file_exists($source) && in_array($pathInfo['extension'], $supportedTypes)) {
	//$ext = $pathInfo['extension'];
	//if (strlen($ext) == 4 && strtolower(substr($ext, -1)) == 'x') {
	//$ext = substr($ext, 0, -1);
	//}
	//
	//$this->source = $source;
	//$this->style = $this->setNewStyle(new ImageStyle(), $style, true);
	//$this->icon = realpath(__DIR__ . "/../resources/{$ext}.png");
}
