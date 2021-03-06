run_mlr -n put -v ''
mlr_expect_fail -n filter -v ''

run_mlr -n put -v 'begin {}'
run_mlr -n put -v 'begin {;}'
run_mlr -n put -v 'begin {;;}'
run_mlr -n put -v 'begin {;;;}'
run_mlr -n put -v 'begin {@x=1}'
run_mlr -n put -v 'begin {@x=1;}'
run_mlr -n put -v 'begin {;@x=1}'
run_mlr -n put -v 'begin {@x=1;@y=2}'
run_mlr -n put -v 'begin {@x=1;;@y=2}'

run_mlr -n put -v 'true {}'
run_mlr -n put -v 'true {;}'
run_mlr -n put -v 'true {;;}'
run_mlr -n put -v 'true {;;;}'
run_mlr -n put -v 'true {@x=1}'
run_mlr -n put -v 'true {@x=1;}'
run_mlr -n put -v 'true {;@x=1}'
run_mlr -n put -v 'true {@x=1;@y=2}'
run_mlr -n put -v 'true {@x=1;;@y=2}'

run_mlr -n put -v 'end {}'
run_mlr -n put -v 'end {;}'
run_mlr -n put -v 'end {;;}'
run_mlr -n put -v 'end {;;;}'
run_mlr -n put -v 'end {@x=1}'
run_mlr -n put -v 'end {@x=1;}'
run_mlr -n put -v 'end {;@x=1}'
run_mlr -n put -v 'end {@x=1;@y=2}'
run_mlr -n put -v 'end {@x=1;;@y=2}'
