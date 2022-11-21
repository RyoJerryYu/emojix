package emojix

import (
	"testing"
)

func TestUnicodeToEmojiTag(t *testing.T) {
	src := `😛👌😇😌👏😒😍😊`
	caseA := UnicodeToEmojiTag(src)
	t.Log(caseA)

	t.Log(EmojiTagToUnicode(caseA))

	if caseA != `:face-with-tongue::ok-hand::smiling-face-with-halo::relieved-face::clapping-hands::unamused-face::smiling-face-with-heart-eyes::smiling-face-with-smiling-eyes:` {
		t.Fatal(`failed to convert emoji chars`, caseA)
	}
}

func TestUnicodeToHTMLEntities(t *testing.T) {
	src := string([]rune{0x30, 0xFE0F, 0x20E3})

	if ret := UnicodeToHTMLEntities(src); ret != `&#x30;&#xFE0F;&#x20E3;` {
		t.Fatal(`failed to convert emoji unicode to html entities.`, ret)
	}
}

func TestUnicodeToHTMLEntities2(t *testing.T) {
	src := string([]rune{0x1f004})

	if ret := UnicodeToHTMLEntities(src); ret != `&#x1F004;` {
		t.Fatal(`failed to convert emoji unicode to html entities.`, ret)
	}

}

func TestEmojiTagToUnicode(t *testing.T) {
	src := `:keyboard:`
	ret := EmojiTagToUnicode(src)

	if ret != string([]rune{0x2328}) {
		t.Fatal(`failed to convert unicode from emoji tag.`, ret)
	}
}

func TestEmojiTagToHTMLEntities(t *testing.T) {
	src := `:keyboard:`
	ret := EmojiTagToHTMLEntities(src)

	if ret != `&#x2328;` {
		t.Fatal(`failed to convert emoji tag to html entities.`, ret)
	}
}

func TestHTMLEntitiesToUnicode(t *testing.T) {
	src := `&#x2328;`
	ret := HTMLEntitiesToUnicode(src)

	if ret != string([]rune{0x2328}) {
		t.Fatal(`failed to html entities to unicode.`, ret)
	}
}
