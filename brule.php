<?php
$all = file('rules.txt');
$rules = array();
foreach($all as $i => $r) {
    if (preg_match('/embed|author|comment|json|trackback/i',$r)) {
        continue;
    }
    if (preg_match('/feed\//i',$r)) {
        continue;
    }
    $r = str_replace('(feed|rdf|rss|rss2|atom)','(feed|rss|atom)',$r);
    $r = str_replace('category_name','name',$r);
    $r = str_replace('?tag=','?name=',$r);
    $r = str_replace('pagename','name',$r);
    $r = str_replace('?&', '', $r);
    $x = explode(' index.php', $r);
    $t = explode('&', $x[1]);
    $x[1] = array();
    foreach($t as $v) {
        $v = explode('=',trim($v,'?'));
        $x[1][] = $v[0];
    }
    //$all[$i] = $x;
    $x[] = $r;
    $rules[] = $x;
}

//foreach($all as $a) {
foreach($rules as $a) {
    $a[0] = str_replace('\\.', '\\\\.', $a[0]);
    echo "        // ".$x[2]."\n        Route{\n            Pattern:\"".$a[0]."\",\n            Params:[]string{\"".implode('","', $a[1])."\"},\n            Controller: DisplayArchiveAllPosts,\n        },\n";
}

//print_r($all);