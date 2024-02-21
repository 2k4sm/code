package Mentor;


public class mentor {
    String Name;
    String Email;
    String Curr_Company;

    public mentor(String name, String email,String company){
        this.Name = name;
        this.Email = email;
        this.Curr_Company = company;
    }

    public mentor(mentor other) {
        this.Name = other.Name;
        this.Email = other.Email;
        this.Curr_Company = other.Curr_Company;
    }

    public String getName(){
        return this.Name;
    }

    public void setName(String Name){
        this.Name = Name;
    }

    public String getEmail(){
        return this.Email;
    }

    public void setEmail(String Email){
        this.Email = Email;
    }

    public String getCompany(){
        return this.Curr_Company;
    }

    public void setCompany(String company){
        this.Curr_Company = company;
    }

}
