#!/bin/zsh
CLASSPATH=.:~/ANTLR-4.13.0/antlr-4.13.0-complete.jar
alias antlr4='java -Xmx500M -cp "~/ANTLR-4.13.0/antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java -Xmx500M -cp "~/ANTLR-4.13.0/antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'

cd grammar

-Dlanguage=Go -o ../chore/parser/ Swift.g4 SwiftLexer.g4 -visitor -no-listener
# antlr4 -Dlanguage=Go -o ../chore/parser/ Swift.g4 -visitor -no-listener

java org.antlr.v4.Tool -Dlanguage=Go -o ../core/parser/ Swift.g4 SwiftLexer.g4 -visitor -no-listener