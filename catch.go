package try

type WrappedError struct {
	err error
}

func (w *WrappedError) Error() string {
	if w.err != nil {
		return w.err.Error()
	}
	return "WrappedError<nil>"
}

func Must(err error) {
	if err != nil {
		panic(WrappedError{err})
	}
}

type Context struct {
	err error
}

func Try(f func()) (ctx Context) {
	defer func() {
		if r := recover(); r != nil {
			if w, ok := r.(WrappedError); ok {
				ctx.err = w.err
			} else {
				panic(r)
			}
		}
	}()
	f()
	return
}

func (ctx Context) Catch(f func(err error)) Context {
	if ctx.err != nil {
		f(ctx.err)
	}
	return ctx
}
