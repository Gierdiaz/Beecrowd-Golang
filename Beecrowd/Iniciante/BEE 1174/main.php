<?php
$vetor = [];

for ($i = 0; $i < 100; $i++) {
    
    $valor = (float) trim(fgets(STDIN));
    
    $vetor[$i] = $valor;

    if ($vetor[$i] <= 10) {
        printf("A[%d] = %.1f\n", $i, $vetor[$i]);
    }
}

?>
