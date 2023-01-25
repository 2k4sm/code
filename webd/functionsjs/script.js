let h1 = document.querySelector("#head")
let add7 = number => (h1.innerHTML = number+7);
function multiply(n1,n2){
    return h1.innerHTML = n1*n2
}
function capitalize(word){
    word = word.toLowerCase()
    word = word.charAt(0).toUpperCase()+word.slice(1)
    
    return h1.innerHTML = word
}
let lastLetter = word => h1.innerHTML = word[word.length-1]



'''
we defined the constants 'EXPECTED_BAKE_TIME' and 'PREPARATION_TIME'
'''

EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

def bake_time_remaining(elapsed_bake_time):
    '''
    This function takes the elapsed_bake_time as argument and return bake time remaining.
    '''
    return EXPECTED_BAKE_TIME - elapsed_bake_time



def preparation_time_in_minutes(layers):
    '''
    This function takes the layers as argument and return preparation time.
    '''
    return PREPARATION_TIME * layers



def elapsed_time_in_minutes(layers,elapsed_bake_time):
    '''
    This function takes layers and  elapsed_bake_time as argument and return elapsed time.
    '''
    return preparation_time_in_minutes(layers)+elapsed_bake_time

