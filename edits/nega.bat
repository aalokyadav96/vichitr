ffmpeg -i %1 -f lavfi -i color=gray -f lavfi -i color=black -f lavfi -i color=white -filter_complex "[0:v]threshold" %1_nega.png