public class Sample implements Runnable{

    private int i;
    private long l;
    private float f;
    private double d;

    @Override
    public void run(){
        f = i + l;
    }

    public static void main(String[] args){
        String str = "Welcome";
        System.out.println(str);
    }
}
