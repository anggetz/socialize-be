PGDMP                      }            smart    17.4    17.4     4           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            5           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            6           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            7           1262    16387    smart    DATABASE     k   CREATE DATABASE smart WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en-US';
    DROP DATABASE smart;
                     postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                     pg_database_owner    false            8           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                        pg_database_owner    false    4            �            1259    16398    topics    TABLE     �   CREATE TABLE public.topics (
    id bigint NOT NULL,
    title character varying,
    content character varying,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    author character varying
);
    DROP TABLE public.topics;
       public         heap r       postgres    false    4            �            1259    16407    topics_comment    TABLE     �   CREATE TABLE public.topics_comment (
    id bigint NOT NULL,
    comment text,
    author character varying,
    created_at timestamp with time zone,
    topics_id bigint
);
 "   DROP TABLE public.topics_comment;
       public         heap r       postgres    false    4            �            1259    16406    topics_comment_id_seq    SEQUENCE     ~   CREATE SEQUENCE public.topics_comment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.topics_comment_id_seq;
       public               postgres    false    220    4            9           0    0    topics_comment_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.topics_comment_id_seq OWNED BY public.topics_comment.id;
          public               postgres    false    219            �            1259    16397    topics_id_seq    SEQUENCE     v   CREATE SEQUENCE public.topics_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.topics_id_seq;
       public               postgres    false    4    218            :           0    0    topics_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.topics_id_seq OWNED BY public.topics.id;
          public               postgres    false    217            �            1259    16421    topics_like    TABLE     �   CREATE TABLE public.topics_like (
    id bigint NOT NULL,
    author character varying,
    topics_id bigint,
    created_at timestamp with time zone
);
    DROP TABLE public.topics_like;
       public         heap r       postgres    false    4            �            1259    16420    topics_like_id_seq    SEQUENCE     {   CREATE SEQUENCE public.topics_like_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.topics_like_id_seq;
       public               postgres    false    222    4            ;           0    0    topics_like_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.topics_like_id_seq OWNED BY public.topics_like.id;
          public               postgres    false    221            �           2604    16401 	   topics id    DEFAULT     f   ALTER TABLE ONLY public.topics ALTER COLUMN id SET DEFAULT nextval('public.topics_id_seq'::regclass);
 8   ALTER TABLE public.topics ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    218    217    218            �           2604    16410    topics_comment id    DEFAULT     v   ALTER TABLE ONLY public.topics_comment ALTER COLUMN id SET DEFAULT nextval('public.topics_comment_id_seq'::regclass);
 @   ALTER TABLE public.topics_comment ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    220    219    220            �           2604    16424    topics_like id    DEFAULT     p   ALTER TABLE ONLY public.topics_like ALTER COLUMN id SET DEFAULT nextval('public.topics_like_id_seq'::regclass);
 =   ALTER TABLE public.topics_like ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    222    221    222            -          0    16398    topics 
   TABLE DATA           T   COPY public.topics (id, title, content, created_at, updated_at, author) FROM stdin;
    public               postgres    false    218          /          0    16407    topics_comment 
   TABLE DATA           T   COPY public.topics_comment (id, comment, author, created_at, topics_id) FROM stdin;
    public               postgres    false    220   �       1          0    16421    topics_like 
   TABLE DATA           H   COPY public.topics_like (id, author, topics_id, created_at) FROM stdin;
    public               postgres    false    222   C        <           0    0    topics_comment_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.topics_comment_id_seq', 39, true);
          public               postgres    false    219            =           0    0    topics_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.topics_id_seq', 8, true);
          public               postgres    false    217            >           0    0    topics_like_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.topics_like_id_seq', 27, true);
          public               postgres    false    221            �           2606    16414 "   topics_comment topics_comment_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.topics_comment
    ADD CONSTRAINT topics_comment_pkey PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.topics_comment DROP CONSTRAINT topics_comment_pkey;
       public                 postgres    false    220            �           2606    16428    topics_like topics_like_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.topics_like
    ADD CONSTRAINT topics_like_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.topics_like DROP CONSTRAINT topics_like_pkey;
       public                 postgres    false    222            �           2606    16405    topics topics_pk 
   CONSTRAINT     N   ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_pk PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.topics DROP CONSTRAINT topics_pk;
       public                 postgres    false    218            �           2606    16415     topics_comment fk_comment_topics    FK CONSTRAINT     �   ALTER TABLE ONLY public.topics_comment
    ADD CONSTRAINT fk_comment_topics FOREIGN KEY (topics_id) REFERENCES public.topics(id) ON DELETE CASCADE;
 J   ALTER TABLE ONLY public.topics_comment DROP CONSTRAINT fk_comment_topics;
       public               postgres    false    220    4756    218            �           2606    16429    topics_like fk_like_topics    FK CONSTRAINT     �   ALTER TABLE ONLY public.topics_like
    ADD CONSTRAINT fk_like_topics FOREIGN KEY (topics_id) REFERENCES public.topics(id) ON DELETE CASCADE;
 D   ALTER TABLE ONLY public.topics_like DROP CONSTRAINT fk_like_topics;
       public               postgres    false    218    222    4756            -   9   x�����N�LJL��4202�50�54R00�25�26�36371�60��"�=... /4
�      /   k   x�}�=�0��9Ev����v���H�H��}�`���$�o'_��������E��)�G,YJ�`�r����T�����тfA2�{�륝���`���s��9�0�'2      1   :   x�32�L��ϫ��/-��4202�50�54R00�25�20�306152�60����� A.
�     