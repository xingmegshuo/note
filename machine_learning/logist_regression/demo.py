
def compression(string):
    data=list(string)
    str_list = [i for i in data if i not in str_list]
    new_str=''
    for i in str_list:
        new_str+=i+data.count(i)
    return new_str


print(compression('aaaabbcddd'))