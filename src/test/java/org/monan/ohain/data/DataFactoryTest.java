package org.monan.ohain.data;

import org.junit.Test;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/23
 */
public class DataFactoryTest {

    @Test
    public void testManu(){
        Example exam = DataFactory.getInstance().manu(Example.class);
    }

}
