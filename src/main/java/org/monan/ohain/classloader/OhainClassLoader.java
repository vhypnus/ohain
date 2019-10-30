package org.monan.ohain.classloader;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/30
 */
public class OhainClassLoader extends ClassLoader {

    public static void main(String[] args) {

        System.out.println(OhainClassLoader.class.getResource("/"));
    }
}
