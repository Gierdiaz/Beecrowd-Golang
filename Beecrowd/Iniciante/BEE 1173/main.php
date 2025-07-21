<?php

$vetor = [];

$valor = (int) trim(fgets(STDIN));

$vetor[0] = $valor;

for ($i = 1; $i < 10; $i++) {   
    $vetor[$i] = $vetor[$i - 1] * 2;
}

for ($i = 0; $i < 10; $i++) {
    echo "N[$i] = {$vetor[$i]}\n";
}

?>
