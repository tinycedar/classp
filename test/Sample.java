public class Sample implements Runnable{
    private String name;

    public String getName(){
        return this.name;
    }

    public void setName(String name){
        this.name = name;
    }

    public void run(){
        
    }

    public static void main(String[] args){
        Sample sample = new Sample();
        sample.setName("Daniel");
        System.out.println("Hello " + sample.getName() + " !");
    }
}
