package main

import (
	"os/exec"
	"log"
)

//~ func ff_to_(fileName , fileExt , desiredExt string) {
	//~ getFrom := "./uploads/" + "/" + fileName + fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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

func ffYelo(f FFFilters) {
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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

func ffVintage(f FFFilters) {
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
	saveAs := "./edits/duotone_" + f.fileName + f.desiredExt

	var filter = "[0:v]format=gray[v0];gradients=s=256x1:c0=0xffffff:c1=0x666666:x0=0:y0=0:x1=256:y1=0,trim=end_frame=1[v1];[v0][v1]paletteuse,format=rgba"
	
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
	getFrom := "./uploads/" + "/" + f.fileName + f.fileExt
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
