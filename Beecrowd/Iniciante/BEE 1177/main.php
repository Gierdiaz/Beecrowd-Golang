<?php

// v = vetor,

$v = [];

$t = (int) trim(fgets(STDIN));

for ($i = 0; $i < 1000; $i++) {
    $v[$i] = $i % $t;
}

for ($i = 0; $i < 1000; $i++) {
    echo "N[$i] = {$v[$i]}\n";
}
