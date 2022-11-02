package fluor

type ParserFunc func(wrappedError error, err *FluentError) *FluentError

// Parser can be overridden to implement custom logic
var Parser ParserFunc = defaultParser

func defaultParser(wrappedError error, err *FluentError) *FluentError {
	return err
}
