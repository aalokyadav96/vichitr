package main

import (
	"os/exec"
	"log"
	"fmt"
	//~ "bytes"
)

//~ func ff_to_(fileName , fileExt , desiredExt string) {
	//~ getFrom := "./uploads/" +fileName + fileExt
	//~ log.Printf(getFrom)
	//~ saveAs := "./edits/" + fileName + desiredExt
	//~ log.Printf(saveAs)
	//~ cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", "scale=-2:480:flags=lanczos", saveAs)
	//~ err := cmd.Start()
	//~ if err != nil {
		//~ log.Fatal(err)
	//~ }
	//~ log.Printf("Waiting for command to finish...")
	//~ err = cmd.Wait()
	//~ log.Printf("Command finished with error: %v", err)
//~ }

func ffScale(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/scale_" + f.fileName + f.desiredExt
	
	var filter = "scale="+f.width+":-2:flags=lanczos"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffNegclr(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/negclr_" + f.fileName + f.desiredExt
	
	var filter = "lutyuv='y=negval',colorchannelmixer=.393:.769:.189:0:.349:.686:.168:0:.272:.534:.131"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffHflip(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/hflip_" + f.fileName + f.desiredExt
	
	var filter = "hflip"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffVflip(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/vflip_" + f.fileName + f.desiredExt
	
	var filter = "vflip"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffGray(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/gray_" + f.fileName + f.desiredExt
	
	var filter = "format=gray"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffRotateClk(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/rotateclk_" + f.fileName + f.desiredExt

	var filter = "transpose=1"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffRotateAngle(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/rotate"+f.angle+"_" + f.fileName + f.desiredExt
	
	//~ var filter = "rotate="+f.angle+"*(PI/180),scale=-2:1080,pad=ih:1080:(ow-iw)/2:(oh-ih)/2"
	//~ var filter = "rotate="+f.angle+"*(PI/180)"
	//~ var filter = "pad=(ih/2+iw):(ih/2+iw):(ow-iw)/2:(oh-ih)/2:color=#00000000,rotate="+f.angle+"*(PI/180)"
	var filter = "format=rgba,pad=(ih/2+iw):(ih/2+iw):(ow-iw)/2:(oh-ih)/2:color=#00000000,rotate="+f.angle+"*(PI/180)"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffRotate180(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/rotate180_" + f.fileName + f.desiredExt
	
	var filter = "vflip,hflip"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffRotateAntiClk(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/rotateanticlk_" + f.fileName + f.desiredExt

	var filter = "transpose=1,transpose=1,transpose=1"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffRemoveChroma(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/remchr_" + f.fileName + f.desiredExt

	var filter = `lutyuv="u=128:v=128"`
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffBurn(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/burn_" + f.fileName + f.desiredExt

	var filter = "lutyuv='y=2*val'"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffBlur(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/blur_" + f.fileName + f.desiredExt

	var filter = "boxblur=10"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffSaturation(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/satu_" + f.fileName + f.desiredExt

	var filter = `lutyuv=u='(val-maxval/2)*2+maxval/2':v='(val-maxval/2)*2+maxval/2'`
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffCanny(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/canny_" + f.fileName + f.desiredExt

	var filter = `edgedetect=low=0:high=0`
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffSepia(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/sepia_" + f.fileName + f.desiredExt

	var filter = `colorchannelmixer=.393:.769:.189:0:.349:.686:.168:0:.272:.534:.131`
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffNegate(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/negate_" + f.fileName + f.desiredExt

	var filter = `"edgedetect=enable='gt(mod(t,60),57)',negate"`
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffNegval(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/negval_" + f.fileName + f.desiredExt

	var filter = "lutrgb='r=negval:g=negval:b=negval'"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffCurveNeg(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/crv_" + f.fileName + f.desiredExt

	var filter = "curves=preset=color_negative"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffYelo(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/yelo_" + f.fileName + f.desiredExt

	var filter = "lutrgb='r=gammaval(0.5):g=val:b=negval'"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffTones(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/tones_" + f.fileName + f.desiredExt

	var filter = `edgedetect=low=0:high=0,negate`
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffEdgeDetect(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/edge_" + f.fileName + f.desiredExt

	var filter = "convolution='1 1 1 1 -8 1 1 1 1:1 1 1 1 -8 1 1 1 1:1 1 1 1 -8 1 1 1 1:1 1 1 1 -8 1 1 1 1:5:5:5:1:0:128:128:0'"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffEmboss(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/emboss_" + f.fileName + f.desiredExt

	var filter = "convolution='-2 -1 0 -1 1 1 0 -1 2:-2 -1 0 -1 1 1 0 -1 2:-2 -1 0 -1 1 1 0 -1 2:-2 -1 0 -1 0 1 0 -1 2'"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffBrr(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/brr_" + f.fileName + f.desiredExt

	var filter = "colorchannelmixer='.964:.456:.865:0:.623:.686:.695:0:.672:.534:.631'"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffVintage(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/vintage_" + f.fileName + f.desiredExt

	var filter = "curves=preset='vintage'"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffDarken(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/darken_" + f.fileName + f.desiredExt

	var filter = "colorlevels=rimin=0.058:gimin=0.058:bimin=0.058"
	
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffDuotone(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/duotone_" + f.fileName + f.desiredExt

	var filter = "[0:v]format=gray[v0];gradients=s=256x1:c0=0xffffff:c1=0x000000:x0=0:y0=0:x1=256:y1=0,trim=end_frame=1[v1];[v0][v1]paletteuse,format=rgba"
	
	log.Println("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffMergePlanes(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/mrgp_" + f.fileName + f.desiredExt

	var filter = "format=yuv444p,mergeplanes=0x000000:yuv444p,format=rgba"
	
	log.Println("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func ffPixelated(f FFFilters) {
	getFrom := "./uploads/" +f.fileName + f.fileExt
	saveAs := "./edits/pixel_" + f.fileName + f.desiredExt

	var filter = "pixelize=w=4:h=4"
	
	log.Println("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ffAlum(f FFFilters) {
	getFrom := "./uploads/" + f.fileName + f.fileExt
	saveAs := "./edits/alum_" + f.fileName + f.desiredExt

	var filter = "selectivecolor=reds='0 0 0 .9':whites='0 0 0 .9'"
	
	log.Println("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func getDims(getFrom string) string {
	//ffprobe -v error -select_streams v -show_entries stream=width,height -of csv=p=0:s=x input.m4v
	out, err := exec.Command("ffprobe", "-v", "error", "-select_streams", "v", "-show_entries", "stream=width,height", "-of", "csv=p=0:s=x", getFrom).Output()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return "0x0"
	}
	return string(out)
}



//~ func ffFix(fixFrom string, f FFFilters) string {
	//~ saveAs := "./edits/dot_" + f.fileName + "_.png"
	//~ dim := getDims(fixFrom)
	//~ fmt.Println(dim)
	
	//~ var filter = "scale=-2:720,pad=720:ih:(ow-iw)/2"
	//~ cmd := exec.Command("ffmpeg", "-i", fixFrom, "-vf", filter, saveAs)
	//~ err := cmd.Start()
	//~ if err != nil {
		//~ log.Fatal(err)
	//~ }
	//~ log.Printf("Waiting for command to finish...")
	//~ err = cmd.Wait()
	//~ log.Printf("Command finished with error: %v", err)
	//~ fmt.Println("./edits/dot_" + f.fileName + "_.png")
	//~ return "./edits/dot_" + f.fileName + "_.png"
//~ }


func ffDotMatrix(f FFFilters) {
	getFrom := "./uploads/" + f.fileName + f.fileExt
	saveAs := "./edits/dot_" + f.fileName + f.desiredExt
	fmt.Println(saveAs)
	dims := ""+f.wd + "x" + f.ht + ""
	//~ dim := getDims(getFrom )
	//~ fmt.Println(dim)
	//~ cmd := exec.Command("ffmpeg", "-i", getFrom, "-f", "lavfi", "-i", "color=gray:s=720x720", "-f", "lavfi", "-i", "color=black:s=720x720", "-f", "lavfi", "-i", "color=white:s=720x720", "-filter_complex", "[0:v]scale=-2:720,pad=720:ih:(ow-iw)/2,threshold", saveAs)
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-f", "lavfi", "-i", "color=gray:s="+dims, "-f", "lavfi", "-i", "color=black:s="+dims, "-f", "lavfi", "-i", "color=white:s="+dims, "-filter_complex", "[0:v]scale=-2:"+f.ht+",threshold", saveAs)
	//~ cmd := exec.Command("ffmpeg", "-i", getFrom, "-f", "lavfi", "-i", "color=0xaaaaaa:s="+dims, "-f", "lavfi", "-i", "color=black:s="+dims, "-f", "lavfi", "-i", "color=white:s="+dim, "-filter_complex", "[0:v]scale="+f.wd+":-2,threshold", saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	//~ var out bytes.Buffer
//~ var stderr bytes.Buffer
//~ cmd.Stdout = &out
//~ cmd.Stderr = &stderr
//~ err := cmd.Run()
//~ if err != nil {
    //~ fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    //~ return
//~ }
//~ fmt.Println("Result: " + out.String())
}


func ffTest(f FFFilters) {
	getFrom := "./uploads/" + f.fileName + f.fileExt
	saveAs := "./edits/test_" + f.fileName + f.desiredExt

	var filter = "colorchannelmixer='.9:.5:.6:0:.6:.3:.6:0:.7:.7:.4'"
	//~ var filter = "colorchannelmixer='1:1:0:0:1:1:0:0:1:1:0:0'"
	//~ var filter = "colorchannelmixer='1:1:1:0:0:1:1:0:0:1:0:0'"
	//~ var filter = "colorchannelmixer='0:0:1:0:0:1:0:0:1:0:0:0'"
	//~ var filter = "colorchannelmixer='1:1:1:0:0:0:1:0:0:1:0:0'"
	//~ var filter = "split=4[a][b][c][d];[b]lutrgb=g=0:b=0[x];[c]lutrgb=r=0:b=0[y];[d]lutrgb=r=0:g=0[z];[a][x][y][z]hstack=4"
	//~ var filter = "colortemperature=temperature=4000"
	//~ var filter = "rgbashift=bv=-50:edge=smear"
	//~ var filter = "rgbashift=rh=-10"
	//~ var filter = "lutyuv=y=negval,curves=preset=negative,colorchannelmixer=.393:.769:.189:0:.349:.686:.168:0:.272:.534:.131"
	//~ var filter = "lutyuv='y=val:u=val:v=val',colorchannelmixer=.964:.456:.865:0:.123:.286:.395:0:.272:.234:.131"
	//~ var filter = "lutrgb='r=1.4*val:g=1.4*val:b=1.4*val',lutyuv='y=val:u=val:v=val'"
	//~ var filter = "convolution='0 -1 0 -1 5 -1 0 -1 0:0 -1 0 -1 5 -1 0 -1 0:0 -1 0 -1 5 -1 0 -1 0:0 -1 0 -1 5 -1 0 -1 0'"
	//~ var filter = "lutyuv=y=negval:u=val:v=val,lutrgb=r=negval:g=negval:b=negval,negate,curves=preset=color_negative"
	//~ var filter = "colorlevels=rimin=0.1:gimin=0.1:bimin=0.1"
	//~ var filter = "dilation"
	//~ var filter = "swapuv"
	//~ var filter = "convolution='1 1 1 1 -8 1 1 1 1:1 1 1 1 -8 1 1 1 1:1 1 1 1 -8 1 1 1 1:1 1 1 1 -8 1 1 1 1:5:5:5:1:0:128:128:0'"
	//~ var filter = "colorchannelmixer='1.2:0:0:0:0:1.1:0:0:0:0:1:0:0:0:0:1'"
	//~ var filter = "colorchannelmixer='.964:.456:.865:0:.623:.686:.695:0:.672:.534:.631'"
	//~ var filter = "chromahold=color=0x000000:blend=0.1"
	//~ var filter = "colorbalance=bh=-0.2:bm=-0.2:bs=-0.2:gh=-0.2:gm=-0.2:gs=-0.2:rh=-0.2:rm=-0.2:rs=-0.2"
	//~ var filter = "lutyuv='y=val:u=val:v=val',colorchannelmixer=.964:.956:.965:0:.523:.586:.595:0:.572:.534:.531"
	//~ var filter = "format=yuv444p,mergeplanes=0x000201:yuv444p"
	//~ var filter = "format=yuv444p,mergeplanes=0x010200:yuv444p"
	//~ var filter = "nlmeans=s=10"
	//~ var filter = "selectivecolor=greens='.5 0 -.33 0':blues='-0 .97'"
	//~ var filter = "shuffleplanes=0:2:1:3"
	//~ var filter = "unsharp=luma_msize_x=7:luma_msize_y=7:luma_amount=2.5"
	//~ var filter = "scale='min(320,iw)':'max(240,ih)'"
	//~ var filter = "lutyuv='y=val:u=val:v=val',colorchannelmixer=.964:.456:.865:0:.423:.686:.395:0:.272:.534:.131"
	
	log.Println("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	cmd := exec.Command("ffmpeg", "-i", getFrom, "-vf", filter, saveAs)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


