import Mentor.mentor;

public class accessmod{
    public static void main(String[] args) {
        mentor newmtr = new mentor("mentor1", "mentor@email", "mentorComp");
        Student newstd = new Student("Sm","sm@mail",newmtr);

        newmtr.setCompany("newcomp");
        System.out.println(newmtr.getCompany());

        System.out.println(newstd.Mentor.getCompany());


    }
}

class Student{
    String Name;
    String Email;
    mentor Mentor;
    private String Password;
    private String PSP;

    public Student(String name, String Email,mentor Mentor){
        this.Name = name;
        this.Email = Email;
        this.Mentor = new mentor(Mentor);
    }

    public void SetPassword(String pssd){
        this.Password = pssd;
    }

    public String getPassword(){
        return this.Password;
    }

    public void SetPsp(String psp){
        this.PSP = psp;
    }


    public String getPsp(){
        return this.PSP;
    }


}