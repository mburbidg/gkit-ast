// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: statement.proto
// Protobuf Java Version: 4.29.3

package ast;

public final class StatementOuterClass {
  private StatementOuterClass() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 29,
      /* patch= */ 3,
      /* suffix= */ "",
      StatementOuterClass.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface StatementOrBuilder extends
      // @@protoc_insertion_point(interface_extends:ast.Statement)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
     * @return Whether the createGraphStatement field is set.
     */
    boolean hasCreateGraphStatement();
    /**
     * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
     * @return The createGraphStatement.
     */
    ast.CatalogStatements.CreateGraphStatement getCreateGraphStatement();
    /**
     * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
     */
    ast.CatalogStatements.CreateGraphStatementOrBuilder getCreateGraphStatementOrBuilder();

    /**
     * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
     * @return Whether the dropGraphStatement field is set.
     */
    boolean hasDropGraphStatement();
    /**
     * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
     * @return The dropGraphStatement.
     */
    ast.CatalogStatements.DropGraphStatement getDropGraphStatement();
    /**
     * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
     */
    ast.CatalogStatements.DropGraphStatementOrBuilder getDropGraphStatementOrBuilder();

    ast.StatementOuterClass.Statement.TypeCase getTypeCase();
  }
  /**
   * Protobuf type {@code ast.Statement}
   */
  public static final class Statement extends
      com.google.protobuf.GeneratedMessage implements
      // @@protoc_insertion_point(message_implements:ast.Statement)
      StatementOrBuilder {
  private static final long serialVersionUID = 0L;
    static {
      com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 29,
        /* patch= */ 3,
        /* suffix= */ "",
        Statement.class.getName());
    }
    // Use Statement.newBuilder() to construct.
    private Statement(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
      super(builder);
    }
    private Statement() {
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return ast.StatementOuterClass.internal_static_ast_Statement_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return ast.StatementOuterClass.internal_static_ast_Statement_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              ast.StatementOuterClass.Statement.class, ast.StatementOuterClass.Statement.Builder.class);
    }

    private int typeCase_ = 0;
    @SuppressWarnings("serial")
    private java.lang.Object type_;
    public enum TypeCase
        implements com.google.protobuf.Internal.EnumLite,
            com.google.protobuf.AbstractMessage.InternalOneOfEnum {
      CREATE_GRAPH_STATEMENT(1),
      DROP_GRAPH_STATEMENT(2),
      TYPE_NOT_SET(0);
      private final int value;
      private TypeCase(int value) {
        this.value = value;
      }
      /**
       * @param value The number of the enum to look for.
       * @return The enum associated with the given number.
       * @deprecated Use {@link #forNumber(int)} instead.
       */
      @java.lang.Deprecated
      public static TypeCase valueOf(int value) {
        return forNumber(value);
      }

      public static TypeCase forNumber(int value) {
        switch (value) {
          case 1: return CREATE_GRAPH_STATEMENT;
          case 2: return DROP_GRAPH_STATEMENT;
          case 0: return TYPE_NOT_SET;
          default: return null;
        }
      }
      public int getNumber() {
        return this.value;
      }
    };

    public TypeCase
    getTypeCase() {
      return TypeCase.forNumber(
          typeCase_);
    }

    public static final int CREATE_GRAPH_STATEMENT_FIELD_NUMBER = 1;
    /**
     * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
     * @return Whether the createGraphStatement field is set.
     */
    @java.lang.Override
    public boolean hasCreateGraphStatement() {
      return typeCase_ == 1;
    }
    /**
     * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
     * @return The createGraphStatement.
     */
    @java.lang.Override
    public ast.CatalogStatements.CreateGraphStatement getCreateGraphStatement() {
      if (typeCase_ == 1) {
         return (ast.CatalogStatements.CreateGraphStatement) type_;
      }
      return ast.CatalogStatements.CreateGraphStatement.getDefaultInstance();
    }
    /**
     * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
     */
    @java.lang.Override
    public ast.CatalogStatements.CreateGraphStatementOrBuilder getCreateGraphStatementOrBuilder() {
      if (typeCase_ == 1) {
         return (ast.CatalogStatements.CreateGraphStatement) type_;
      }
      return ast.CatalogStatements.CreateGraphStatement.getDefaultInstance();
    }

    public static final int DROP_GRAPH_STATEMENT_FIELD_NUMBER = 2;
    /**
     * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
     * @return Whether the dropGraphStatement field is set.
     */
    @java.lang.Override
    public boolean hasDropGraphStatement() {
      return typeCase_ == 2;
    }
    /**
     * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
     * @return The dropGraphStatement.
     */
    @java.lang.Override
    public ast.CatalogStatements.DropGraphStatement getDropGraphStatement() {
      if (typeCase_ == 2) {
         return (ast.CatalogStatements.DropGraphStatement) type_;
      }
      return ast.CatalogStatements.DropGraphStatement.getDefaultInstance();
    }
    /**
     * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
     */
    @java.lang.Override
    public ast.CatalogStatements.DropGraphStatementOrBuilder getDropGraphStatementOrBuilder() {
      if (typeCase_ == 2) {
         return (ast.CatalogStatements.DropGraphStatement) type_;
      }
      return ast.CatalogStatements.DropGraphStatement.getDefaultInstance();
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (typeCase_ == 1) {
        output.writeMessage(1, (ast.CatalogStatements.CreateGraphStatement) type_);
      }
      if (typeCase_ == 2) {
        output.writeMessage(2, (ast.CatalogStatements.DropGraphStatement) type_);
      }
      getUnknownFields().writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (typeCase_ == 1) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, (ast.CatalogStatements.CreateGraphStatement) type_);
      }
      if (typeCase_ == 2) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(2, (ast.CatalogStatements.DropGraphStatement) type_);
      }
      size += getUnknownFields().getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof ast.StatementOuterClass.Statement)) {
        return super.equals(obj);
      }
      ast.StatementOuterClass.Statement other = (ast.StatementOuterClass.Statement) obj;

      if (!getTypeCase().equals(other.getTypeCase())) return false;
      switch (typeCase_) {
        case 1:
          if (!getCreateGraphStatement()
              .equals(other.getCreateGraphStatement())) return false;
          break;
        case 2:
          if (!getDropGraphStatement()
              .equals(other.getDropGraphStatement())) return false;
          break;
        case 0:
        default:
      }
      if (!getUnknownFields().equals(other.getUnknownFields())) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      switch (typeCase_) {
        case 1:
          hash = (37 * hash) + CREATE_GRAPH_STATEMENT_FIELD_NUMBER;
          hash = (53 * hash) + getCreateGraphStatement().hashCode();
          break;
        case 2:
          hash = (37 * hash) + DROP_GRAPH_STATEMENT_FIELD_NUMBER;
          hash = (53 * hash) + getDropGraphStatement().hashCode();
          break;
        case 0:
        default:
      }
      hash = (29 * hash) + getUnknownFields().hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static ast.StatementOuterClass.Statement parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static ast.StatementOuterClass.Statement parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static ast.StatementOuterClass.Statement parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static ast.StatementOuterClass.Statement parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static ast.StatementOuterClass.Statement parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static ast.StatementOuterClass.Statement parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static ast.StatementOuterClass.Statement parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static ast.StatementOuterClass.Statement parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static ast.StatementOuterClass.Statement parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static ast.StatementOuterClass.Statement parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static ast.StatementOuterClass.Statement parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static ast.StatementOuterClass.Statement parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(ast.StatementOuterClass.Statement prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessage.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code ast.Statement}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessage.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:ast.Statement)
        ast.StatementOuterClass.StatementOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return ast.StatementOuterClass.internal_static_ast_Statement_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return ast.StatementOuterClass.internal_static_ast_Statement_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                ast.StatementOuterClass.Statement.class, ast.StatementOuterClass.Statement.Builder.class);
      }

      // Construct using ast.StatementOuterClass.Statement.newBuilder()
      private Builder() {

      }

      private Builder(
          com.google.protobuf.GeneratedMessage.BuilderParent parent) {
        super(parent);

      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        bitField0_ = 0;
        if (createGraphStatementBuilder_ != null) {
          createGraphStatementBuilder_.clear();
        }
        if (dropGraphStatementBuilder_ != null) {
          dropGraphStatementBuilder_.clear();
        }
        typeCase_ = 0;
        type_ = null;
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return ast.StatementOuterClass.internal_static_ast_Statement_descriptor;
      }

      @java.lang.Override
      public ast.StatementOuterClass.Statement getDefaultInstanceForType() {
        return ast.StatementOuterClass.Statement.getDefaultInstance();
      }

      @java.lang.Override
      public ast.StatementOuterClass.Statement build() {
        ast.StatementOuterClass.Statement result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public ast.StatementOuterClass.Statement buildPartial() {
        ast.StatementOuterClass.Statement result = new ast.StatementOuterClass.Statement(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        buildPartialOneofs(result);
        onBuilt();
        return result;
      }

      private void buildPartial0(ast.StatementOuterClass.Statement result) {
        int from_bitField0_ = bitField0_;
      }

      private void buildPartialOneofs(ast.StatementOuterClass.Statement result) {
        result.typeCase_ = typeCase_;
        result.type_ = this.type_;
        if (typeCase_ == 1 &&
            createGraphStatementBuilder_ != null) {
          result.type_ = createGraphStatementBuilder_.build();
        }
        if (typeCase_ == 2 &&
            dropGraphStatementBuilder_ != null) {
          result.type_ = dropGraphStatementBuilder_.build();
        }
      }

      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof ast.StatementOuterClass.Statement) {
          return mergeFrom((ast.StatementOuterClass.Statement)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(ast.StatementOuterClass.Statement other) {
        if (other == ast.StatementOuterClass.Statement.getDefaultInstance()) return this;
        switch (other.getTypeCase()) {
          case CREATE_GRAPH_STATEMENT: {
            mergeCreateGraphStatement(other.getCreateGraphStatement());
            break;
          }
          case DROP_GRAPH_STATEMENT: {
            mergeDropGraphStatement(other.getDropGraphStatement());
            break;
          }
          case TYPE_NOT_SET: {
            break;
          }
        }
        this.mergeUnknownFields(other.getUnknownFields());
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        if (extensionRegistry == null) {
          throw new java.lang.NullPointerException();
        }
        try {
          boolean done = false;
          while (!done) {
            int tag = input.readTag();
            switch (tag) {
              case 0:
                done = true;
                break;
              case 10: {
                input.readMessage(
                    getCreateGraphStatementFieldBuilder().getBuilder(),
                    extensionRegistry);
                typeCase_ = 1;
                break;
              } // case 10
              case 18: {
                input.readMessage(
                    getDropGraphStatementFieldBuilder().getBuilder(),
                    extensionRegistry);
                typeCase_ = 2;
                break;
              } // case 18
              default: {
                if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                  done = true; // was an endgroup tag
                }
                break;
              } // default:
            } // switch (tag)
          } // while (!done)
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.unwrapIOException();
        } finally {
          onChanged();
        } // finally
        return this;
      }
      private int typeCase_ = 0;
      private java.lang.Object type_;
      public TypeCase
          getTypeCase() {
        return TypeCase.forNumber(
            typeCase_);
      }

      public Builder clearType() {
        typeCase_ = 0;
        type_ = null;
        onChanged();
        return this;
      }

      private int bitField0_;

      private com.google.protobuf.SingleFieldBuilder<
          ast.CatalogStatements.CreateGraphStatement, ast.CatalogStatements.CreateGraphStatement.Builder, ast.CatalogStatements.CreateGraphStatementOrBuilder> createGraphStatementBuilder_;
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       * @return Whether the createGraphStatement field is set.
       */
      @java.lang.Override
      public boolean hasCreateGraphStatement() {
        return typeCase_ == 1;
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       * @return The createGraphStatement.
       */
      @java.lang.Override
      public ast.CatalogStatements.CreateGraphStatement getCreateGraphStatement() {
        if (createGraphStatementBuilder_ == null) {
          if (typeCase_ == 1) {
            return (ast.CatalogStatements.CreateGraphStatement) type_;
          }
          return ast.CatalogStatements.CreateGraphStatement.getDefaultInstance();
        } else {
          if (typeCase_ == 1) {
            return createGraphStatementBuilder_.getMessage();
          }
          return ast.CatalogStatements.CreateGraphStatement.getDefaultInstance();
        }
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       */
      public Builder setCreateGraphStatement(ast.CatalogStatements.CreateGraphStatement value) {
        if (createGraphStatementBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          type_ = value;
          onChanged();
        } else {
          createGraphStatementBuilder_.setMessage(value);
        }
        typeCase_ = 1;
        return this;
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       */
      public Builder setCreateGraphStatement(
          ast.CatalogStatements.CreateGraphStatement.Builder builderForValue) {
        if (createGraphStatementBuilder_ == null) {
          type_ = builderForValue.build();
          onChanged();
        } else {
          createGraphStatementBuilder_.setMessage(builderForValue.build());
        }
        typeCase_ = 1;
        return this;
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       */
      public Builder mergeCreateGraphStatement(ast.CatalogStatements.CreateGraphStatement value) {
        if (createGraphStatementBuilder_ == null) {
          if (typeCase_ == 1 &&
              type_ != ast.CatalogStatements.CreateGraphStatement.getDefaultInstance()) {
            type_ = ast.CatalogStatements.CreateGraphStatement.newBuilder((ast.CatalogStatements.CreateGraphStatement) type_)
                .mergeFrom(value).buildPartial();
          } else {
            type_ = value;
          }
          onChanged();
        } else {
          if (typeCase_ == 1) {
            createGraphStatementBuilder_.mergeFrom(value);
          } else {
            createGraphStatementBuilder_.setMessage(value);
          }
        }
        typeCase_ = 1;
        return this;
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       */
      public Builder clearCreateGraphStatement() {
        if (createGraphStatementBuilder_ == null) {
          if (typeCase_ == 1) {
            typeCase_ = 0;
            type_ = null;
            onChanged();
          }
        } else {
          if (typeCase_ == 1) {
            typeCase_ = 0;
            type_ = null;
          }
          createGraphStatementBuilder_.clear();
        }
        return this;
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       */
      public ast.CatalogStatements.CreateGraphStatement.Builder getCreateGraphStatementBuilder() {
        return getCreateGraphStatementFieldBuilder().getBuilder();
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       */
      @java.lang.Override
      public ast.CatalogStatements.CreateGraphStatementOrBuilder getCreateGraphStatementOrBuilder() {
        if ((typeCase_ == 1) && (createGraphStatementBuilder_ != null)) {
          return createGraphStatementBuilder_.getMessageOrBuilder();
        } else {
          if (typeCase_ == 1) {
            return (ast.CatalogStatements.CreateGraphStatement) type_;
          }
          return ast.CatalogStatements.CreateGraphStatement.getDefaultInstance();
        }
      }
      /**
       * <code>.ast.CreateGraphStatement create_graph_statement = 1;</code>
       */
      private com.google.protobuf.SingleFieldBuilder<
          ast.CatalogStatements.CreateGraphStatement, ast.CatalogStatements.CreateGraphStatement.Builder, ast.CatalogStatements.CreateGraphStatementOrBuilder> 
          getCreateGraphStatementFieldBuilder() {
        if (createGraphStatementBuilder_ == null) {
          if (!(typeCase_ == 1)) {
            type_ = ast.CatalogStatements.CreateGraphStatement.getDefaultInstance();
          }
          createGraphStatementBuilder_ = new com.google.protobuf.SingleFieldBuilder<
              ast.CatalogStatements.CreateGraphStatement, ast.CatalogStatements.CreateGraphStatement.Builder, ast.CatalogStatements.CreateGraphStatementOrBuilder>(
                  (ast.CatalogStatements.CreateGraphStatement) type_,
                  getParentForChildren(),
                  isClean());
          type_ = null;
        }
        typeCase_ = 1;
        onChanged();
        return createGraphStatementBuilder_;
      }

      private com.google.protobuf.SingleFieldBuilder<
          ast.CatalogStatements.DropGraphStatement, ast.CatalogStatements.DropGraphStatement.Builder, ast.CatalogStatements.DropGraphStatementOrBuilder> dropGraphStatementBuilder_;
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       * @return Whether the dropGraphStatement field is set.
       */
      @java.lang.Override
      public boolean hasDropGraphStatement() {
        return typeCase_ == 2;
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       * @return The dropGraphStatement.
       */
      @java.lang.Override
      public ast.CatalogStatements.DropGraphStatement getDropGraphStatement() {
        if (dropGraphStatementBuilder_ == null) {
          if (typeCase_ == 2) {
            return (ast.CatalogStatements.DropGraphStatement) type_;
          }
          return ast.CatalogStatements.DropGraphStatement.getDefaultInstance();
        } else {
          if (typeCase_ == 2) {
            return dropGraphStatementBuilder_.getMessage();
          }
          return ast.CatalogStatements.DropGraphStatement.getDefaultInstance();
        }
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       */
      public Builder setDropGraphStatement(ast.CatalogStatements.DropGraphStatement value) {
        if (dropGraphStatementBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          type_ = value;
          onChanged();
        } else {
          dropGraphStatementBuilder_.setMessage(value);
        }
        typeCase_ = 2;
        return this;
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       */
      public Builder setDropGraphStatement(
          ast.CatalogStatements.DropGraphStatement.Builder builderForValue) {
        if (dropGraphStatementBuilder_ == null) {
          type_ = builderForValue.build();
          onChanged();
        } else {
          dropGraphStatementBuilder_.setMessage(builderForValue.build());
        }
        typeCase_ = 2;
        return this;
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       */
      public Builder mergeDropGraphStatement(ast.CatalogStatements.DropGraphStatement value) {
        if (dropGraphStatementBuilder_ == null) {
          if (typeCase_ == 2 &&
              type_ != ast.CatalogStatements.DropGraphStatement.getDefaultInstance()) {
            type_ = ast.CatalogStatements.DropGraphStatement.newBuilder((ast.CatalogStatements.DropGraphStatement) type_)
                .mergeFrom(value).buildPartial();
          } else {
            type_ = value;
          }
          onChanged();
        } else {
          if (typeCase_ == 2) {
            dropGraphStatementBuilder_.mergeFrom(value);
          } else {
            dropGraphStatementBuilder_.setMessage(value);
          }
        }
        typeCase_ = 2;
        return this;
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       */
      public Builder clearDropGraphStatement() {
        if (dropGraphStatementBuilder_ == null) {
          if (typeCase_ == 2) {
            typeCase_ = 0;
            type_ = null;
            onChanged();
          }
        } else {
          if (typeCase_ == 2) {
            typeCase_ = 0;
            type_ = null;
          }
          dropGraphStatementBuilder_.clear();
        }
        return this;
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       */
      public ast.CatalogStatements.DropGraphStatement.Builder getDropGraphStatementBuilder() {
        return getDropGraphStatementFieldBuilder().getBuilder();
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       */
      @java.lang.Override
      public ast.CatalogStatements.DropGraphStatementOrBuilder getDropGraphStatementOrBuilder() {
        if ((typeCase_ == 2) && (dropGraphStatementBuilder_ != null)) {
          return dropGraphStatementBuilder_.getMessageOrBuilder();
        } else {
          if (typeCase_ == 2) {
            return (ast.CatalogStatements.DropGraphStatement) type_;
          }
          return ast.CatalogStatements.DropGraphStatement.getDefaultInstance();
        }
      }
      /**
       * <code>.ast.DropGraphStatement drop_graph_statement = 2;</code>
       */
      private com.google.protobuf.SingleFieldBuilder<
          ast.CatalogStatements.DropGraphStatement, ast.CatalogStatements.DropGraphStatement.Builder, ast.CatalogStatements.DropGraphStatementOrBuilder> 
          getDropGraphStatementFieldBuilder() {
        if (dropGraphStatementBuilder_ == null) {
          if (!(typeCase_ == 2)) {
            type_ = ast.CatalogStatements.DropGraphStatement.getDefaultInstance();
          }
          dropGraphStatementBuilder_ = new com.google.protobuf.SingleFieldBuilder<
              ast.CatalogStatements.DropGraphStatement, ast.CatalogStatements.DropGraphStatement.Builder, ast.CatalogStatements.DropGraphStatementOrBuilder>(
                  (ast.CatalogStatements.DropGraphStatement) type_,
                  getParentForChildren(),
                  isClean());
          type_ = null;
        }
        typeCase_ = 2;
        onChanged();
        return dropGraphStatementBuilder_;
      }

      // @@protoc_insertion_point(builder_scope:ast.Statement)
    }

    // @@protoc_insertion_point(class_scope:ast.Statement)
    private static final ast.StatementOuterClass.Statement DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new ast.StatementOuterClass.Statement();
    }

    public static ast.StatementOuterClass.Statement getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Statement>
        PARSER = new com.google.protobuf.AbstractParser<Statement>() {
      @java.lang.Override
      public Statement parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        Builder builder = newBuilder();
        try {
          builder.mergeFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.setUnfinishedMessage(builder.buildPartial());
        } catch (com.google.protobuf.UninitializedMessageException e) {
          throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
        } catch (java.io.IOException e) {
          throw new com.google.protobuf.InvalidProtocolBufferException(e)
              .setUnfinishedMessage(builder.buildPartial());
        }
        return builder.buildPartial();
      }
    };

    public static com.google.protobuf.Parser<Statement> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Statement> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public ast.StatementOuterClass.Statement getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_ast_Statement_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_ast_Statement_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\017statement.proto\022\003ast\032!google/protobuf/" +
      "go_features.proto\032\030catalog_statements.pr" +
      "oto\"\211\001\n\tStatement\022;\n\026create_graph_statem" +
      "ent\030\001 \001(\0132\031.ast.CreateGraphStatementH\000\0227" +
      "\n\024drop_graph_statement\030\002 \001(\0132\027.ast.DropG" +
      "raphStatementH\000B\006\n\004typeB\016Z\004/ast\222\003\005\322>\002\020\003b" +
      "\010editionsp\350\007"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          pb.GoFeaturesOuterClass.getDescriptor(),
          ast.CatalogStatements.getDescriptor(),
        });
    internal_static_ast_Statement_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_ast_Statement_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_ast_Statement_descriptor,
        new java.lang.String[] { "CreateGraphStatement", "DropGraphStatement", "Type", });
    descriptor.resolveAllFeaturesImmutable();
    pb.GoFeaturesOuterClass.getDescriptor();
    ast.CatalogStatements.getDescriptor();
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(pb.GoFeaturesOuterClass.go);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
  }

  // @@protoc_insertion_point(outer_class_scope)
}
