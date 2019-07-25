<?php
$all = file('rules.txt');
$rules = array();
foreach($all as $i => $r) {
    if (preg_match('/embed|comment|json|trackback/i',$r)) {
        continue;
    }
    if (preg_match('/feed\//i',$r)) {
        continue;
    }
    $aa = $r;
    $ctl = 'DisplayHomePageHandle';
    if (preg_match('/\?year=/i',$r)) {
        $ctl = 'DisplayArchiveHandle';
    } else if (preg_match('/\?pagename=/',$r)) {
        $ctl = 'DisplayPageHandle';
    } else if (preg_match('/\?name=/',$r)) {
        $ctl = 'DisplayPostHandle';
    } else if (preg_match('/\?tag=/',$r)) {
        $ctl = 'DisplayTagHandle';
    } else if (preg_match('/\?category_name=/',$r)) {
        $ctl = 'DisplayCategoryHandle';
    } else if (preg_match('/sitemap/',$r)) {
        $ctl = 'DisplaySiteMapHandle';
    } else if (preg_match('/search\//',$r)) {
        $ctl = 'DisplaySearchHandle';
    } else if (preg_match('/author\//',$r)) {
        $ctl = 'DisplayAuthorHandle';
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
    $x[] = $aa;
    $x[] = $ctl;
    $rules[] = $x;
}

//foreach($all as $a) {
foreach($rules as $a) {
    $a[0] = str_replace('\\.', '\\\\.', $a[0]);
    // echo "        // ".trim($a[2])."\n        Route{\n            Pattern:\"".$a[0]."\",\n            Params:[]string{\"".implode('","', $a[1])."\"},\n            Controller: ".$a[3].",\n        },\n";
    echo "        Route{\n            Pattern:\"".$a[0]."\",\n            Params:[]string{\"".implode('","', $a[1])."\"},\n            Controller: ".$a[3].",\n        },\n";
}

//print_r($all);
