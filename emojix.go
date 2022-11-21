package emojix

import (
	"fmt"
	"strings"

	"github.com/forPelevin/gomoji"
)

var (
	unicode2HTMLEntity  *strings.Replacer
	htmlEntity2Unicode  *strings.Replacer
	emojiTag2HTMLEntity *strings.Replacer
	htmlEntity2EmojiTag *strings.Replacer
	unicode2EmojiTag    *strings.Replacer
	emojiTag2Unicode    *strings.Replacer
)

func mapToReplacer(m map[string]string) *strings.Replacer {
	var keys []string
	for fromStr, toStr := range m {
		keys = append(keys, fromStr, toStr)
	}

	return strings.NewReplacer(keys...)
}

func init() {
	allEmojis := gomoji.AllEmojis()

	funcToReplacer := func(
		fIn func(gomoji.Emoji) string,
		fOut func(gomoji.Emoji) string,
	) *strings.Replacer {
		m := make(map[string]string)
		for _, em := range allEmojis {
			m[fIn(em)] = fOut(em)
		}

		return mapToReplacer(m)
	}

	unicode2HTMLEntity = funcToReplacer(emojiToUnicode, emojiToHTMLEntity)
	htmlEntity2Unicode = funcToReplacer(emojiToHTMLEntity, emojiToUnicode)
	emojiTag2HTMLEntity = funcToReplacer(emojiToEmojiTag, emojiToHTMLEntity)
	htmlEntity2EmojiTag = funcToReplacer(emojiToHTMLEntity, emojiToEmojiTag)
	unicode2EmojiTag = funcToReplacer(emojiToUnicode, emojiToEmojiTag)
	emojiTag2Unicode = funcToReplacer(emojiToEmojiTag, emojiToUnicode)
}

func code2entities(src []rune) string {
	ret := make([]string, 0)

	for _, char := range src {
		ret = append(ret, fmt.Sprintf(`&#x%X;`, char))
	}

	return strings.Join(ret, ``)
}

func emojiToUnicode(em gomoji.Emoji) string {
	return em.Character
}

func emojiToHTMLEntity(em gomoji.Emoji) string {
	return code2entities([]rune(em.Character))
}

func emojiToEmojiTag(em gomoji.Emoji) string {
	return fmt.Sprintf(":%s:", em.Slug)
}

// EmojiTagToHTMLEntities is replacing emoji tag to html entities of emoji unicode.
//
// For example: :+1: => &#x1F44D;
func EmojiTagToHTMLEntities(src string) string {
	return emojiTag2HTMLEntity.Replace(src)
}

// HTMLEntitiesToEmojiTag is replacing html entities of emoji unicode to emoji tag.
//
// For Example: &#x1F44D; => :+1:
func HTMLEntitiesToEmojiTag(src string) string {
	return htmlEntity2EmojiTag.Replace(src)
}

// UnicodeToEmojiTag is replacing emoji tag to unicode chars.
//
// For example: rune(0x1F44D) => :+1:
func UnicodeToEmojiTag(src string) string {
	return unicode2EmojiTag.Replace(src)
}

// EmojiTagToUnicode is replacing emoji tag to unicode chars.
//
// For example: :+1: => rune(0x1F44D)
func EmojiTagToUnicode(src string) string {
	return emojiTag2Unicode.Replace(src)
}

// UnicodeToHTMLEntities is replacing unicode emoji chars to html entities.
//
// For Example: rune(0x1F44D) => &#x1F44D;
func UnicodeToHTMLEntities(src string) string {
	return unicode2HTMLEntity.Replace(src)
}

// HTMLEntitiesToUnicode is replacing html entities of emoji unicode to unicode chars.
//
// For Example: &#x1F44D; => rune(0x1F44D)
func HTMLEntitiesToUnicode(src string) string {
	return htmlEntity2Unicode.Replace(src)
}
