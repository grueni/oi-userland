(cd otp_doc_man_22.2/; gfind *  -mindepth 2 -type f |nawk '{print "file path=usr/share/" $1}')|gsed  's/man3/man3erl/'|gsed  's/man4/man5erl/'|gsed  's/man7/man7erl/'|gsed  's/man6/man7erl/' > man0.log
gsed 's/.6$/.7erl/' man0.log|gsed 's/.3$/.3erl/'|gsed 's/.4$/.5erl/'|gsed 's/.7$/.7erl/' > man.log
(cd otp_doc_html_22.2/; gfind *  -mindepth 2 -type f |nawk '{print "file path=usr/share/doc/erlang/" $1}') > html.log

