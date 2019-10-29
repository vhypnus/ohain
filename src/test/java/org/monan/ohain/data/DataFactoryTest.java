package org.monan.ohain.data;

import com.github.javafaker.Faker;
import org.junit.Test;
import org.monan.ohain.data.factory.DataFactory;

import java.util.Locale;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/23
 */
public class DataFactoryTest {

    @Test
    public void testManu(){
//        Example exam = DataFactory.getInstance().manu(Example.class);
        Faker faker = new Faker(new Locale("zh-CN")) ;
        System.out.println(faker.company().name()) ;
//        faker.number().
    }


}
