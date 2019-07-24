<?php
$all = file('rules.txt');

foreach($all as $i => $r) {
    $r = str_replace('?&', '', $r);
    $x = explode(' index.php', $r);
    $t = explode('&', $x[1]);
    $x[1] = array();
    foreach($t as $v) {
        $v = explode('=',trim($v,'?'));
        $x[1][] = $v[0];
    }
    $all[$i] = $x;
}

foreach($all as $a) {
    $a[0] = str_replace('\\.', '\\\\.', $a[0]);
    echo "        Route{\n            Pattern:\"".$a[0]."\",\n            Params:[]string{\"".implode('","', $a[1])."\"},\n            Controller: DisplayArchiveAllPosts,\n        },\n";
}

//print_r($all);