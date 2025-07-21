<?php

$vetores = [];

for ($i = 0; $i < 10; $i++) {
    $valores = (int) trim(fgets(STDIN));

    if ($valores <= 0) {
        $valores = 1;
    }

    $vetores[$i] = $valores;

    echo "X[$i] = {$vetores[$i]}\n";
}
?>
