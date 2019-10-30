package org.monan.ohain.proxy;

import org.junit.Assert;
import org.junit.Test;
import org.monan.ohain.example.service.impl.ExampleServiceImpl;

public class ProxyTest {

    @Test
    public void test_get(){
        Proxy proxy = new Proxy() ;
        Class clazz = proxy.get(ExampleServiceImpl.class) ;
        Assert.assertNotNull(clazz);
    }
}
