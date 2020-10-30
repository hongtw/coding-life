def substring_test(s1, s2):
    s1 = s1.lower()
    s2 = s2.lower()
    for i in range(len(s1)-1):
        if s1[i:i+2] in s2:
            return True
    return False



def test(func):

    assert func("Something","Home") == True
    assert func("Something","Fun") == False
    assert func("Something","") == False
    assert func("","Something") == False
    assert func("BANANA","banana") == True
    assert func("test","lllt") == False
    assert func("","") == False
    assert func("1234567","541265") == True
    assert func("supercalifragilisticexpialidocious","SoundOfItIsAtrocious") == True
    assert func("LoremipsumdolorsitametconsecteturadipiscingelitAeneannonaliquetligulautplaceratorciSuspendissepotentiMorbivolutpatauctoripsumegetaliquamPhasellusidmagnaelitNullamerostellustemporquismolestieaornarevitaediamNullaaliquamrisusnonviverrasagittisInlaoreetultricespretiumVestibulumegetnullatinciduntsempersemacrutrumfelisPraesentpurusarcutempusnecvariusidultricesaduiPellentesqueultriciesjustolobortisrhoncusdignissimNuncviverraconsequatblanditUtbibendumatlacusactristiqueAliquamimperdietnuncsempertortorefficiturviverra","thisisalongstringtest") == True
    print(f"[{func.__name__}] All test is complete")

test(substring_test)