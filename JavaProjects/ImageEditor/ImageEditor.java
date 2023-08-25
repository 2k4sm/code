package ImageEditor;
import java.awt.image.*;

public class ImageEditor{
    public static void main(String[] args) {
        
    }

    
    //Write a function to mirror the image.
    public static BufferedImage mirrorImage(BufferedImage image){
        int imageHeight = image.getHeight();
        int imageWidth = image.getWidth();
        BufferedImage outimage = new BufferedImage(imageWidth,imageHeight,BufferedImage.TYPE_3BYTE_BGR);
        for(int i = 0; i<imageHeight;i++){
            for(int j = 0; j<imageWidth/2;j++){
                outimage.setRGB(imageWidth-1-j,i,image.getRGB(j,i));                
                outimage.setRGB(j,i,image.getRGB(imageWidth-1-j,i));
            }
        }
        return outimage;
    }

    //Write a Function to Rotate the image by 90 degrees.
    public static BufferedImage rotateImage(BufferedImage image){
        int imageHeight = image.getHeight();
        int imageWidth = image.getWidth();

        BufferedImage outImage = new BufferedImage(imageHeight,imageWidth,BufferedImage.TYPE_3BYTE_BGR);

        //Write a loop to transpose the image.
        for(int i = 0; i<imageWidth;i++){
            for(int j = 0; j<imageHeight;j++){
                outImage.setRGB(j, i,image.getRGB(i,j));
            }
        }
        outImage = mirrorImage(outImage);

        return outImage;
    }

    //Write a function to return a grayScale Image.
    public static BufferedImage grayScale(BufferedImage image){
        int imageWidth = image.getWidth();
        int imageHeight = image.getHeight();

    }

    //Brightness control.
}