from ast import NotIn


def count_ln():
    file = open("./testfile.txt","r+")
    
    line_cnt = 0
    try:
        while True:
            i= file.readline()
            if i[0] not in "aeiouAEIOU":
                line_cnt+=1
                print("this is line: ",line_cnt)
                print(i)
            
            else:
                line_cnt+=1
    except IndexError:
        print("Task finished")
count_ln()