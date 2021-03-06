run_mlr repl <<EOF
EOF

run_mlr repl <<EOF
x=1
y=2
print x+y
EOF

run_mlr repl <<EOF
\$x=3
\$*
EOF

run_mlr repl <<EOF
<
begin {
  print "In the beginning:"
}
end {
  print "At the end:"
}
# Populates the main block
print "In ...";
print "... the middle!"
>

begin { print "HELLO" }
begin { print "WORLD" }
end   { print "GOODBYE" }
end   { print "WORLD" }

# Immediately executed
print "HOW ARE THINGS?"

:blocks

:begin
:main
:end
EOF
