// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: program.proto
// Protobuf Java Version: 4.29.3

package ast;

public final class ProgramOuterClass {
  private ProgramOuterClass() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 29,
      /* patch= */ 3,
      /* suffix= */ "",
      ProgramOuterClass.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ProgramOrBuilder extends
      // @@protoc_insertion_point(interface_extends:ast.Program)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>.ast.ProgramActivity program_activity = 1;</code>
     * @return Whether the programActivity field is set.
     */
    boolean hasProgramActivity();
    /**
     * <code>.ast.ProgramActivity program_activity = 1;</code>
     * @return The programActivity.
     */
    ast.ProgramActivityOuterClass.ProgramActivity getProgramActivity();
    /**
     * <code>.ast.ProgramActivity program_activity = 1;</code>
     */
    ast.ProgramActivityOuterClass.ProgramActivityOrBuilder getProgramActivityOrBuilder();

    /**
     * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
     * @return Whether the sessionCloseCommand field is set.
     */
    boolean hasSessionCloseCommand();
    /**
     * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
     * @return The sessionCloseCommand.
     */
    ast.SessionCommands.SessionCloseCommand getSessionCloseCommand();
    /**
     * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
     */
    ast.SessionCommands.SessionCloseCommandOrBuilder getSessionCloseCommandOrBuilder();
  }
  /**
   * Protobuf type {@code ast.Program}
   */
  public static final class Program extends
      com.google.protobuf.GeneratedMessage implements
      // @@protoc_insertion_point(message_implements:ast.Program)
      ProgramOrBuilder {
  private static final long serialVersionUID = 0L;
    static {
      com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 29,
        /* patch= */ 3,
        /* suffix= */ "",
        Program.class.getName());
    }
    // Use Program.newBuilder() to construct.
    private Program(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
      super(builder);
    }
    private Program() {
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return ast.ProgramOuterClass.internal_static_ast_Program_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return ast.ProgramOuterClass.internal_static_ast_Program_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              ast.ProgramOuterClass.Program.class, ast.ProgramOuterClass.Program.Builder.class);
    }

    private int bitField0_;
    public static final int PROGRAM_ACTIVITY_FIELD_NUMBER = 1;
    private ast.ProgramActivityOuterClass.ProgramActivity programActivity_;
    /**
     * <code>.ast.ProgramActivity program_activity = 1;</code>
     * @return Whether the programActivity field is set.
     */
    @java.lang.Override
    public boolean hasProgramActivity() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <code>.ast.ProgramActivity program_activity = 1;</code>
     * @return The programActivity.
     */
    @java.lang.Override
    public ast.ProgramActivityOuterClass.ProgramActivity getProgramActivity() {
      return programActivity_ == null ? ast.ProgramActivityOuterClass.ProgramActivity.getDefaultInstance() : programActivity_;
    }
    /**
     * <code>.ast.ProgramActivity program_activity = 1;</code>
     */
    @java.lang.Override
    public ast.ProgramActivityOuterClass.ProgramActivityOrBuilder getProgramActivityOrBuilder() {
      return programActivity_ == null ? ast.ProgramActivityOuterClass.ProgramActivity.getDefaultInstance() : programActivity_;
    }

    public static final int SESSION_CLOSE_COMMAND_FIELD_NUMBER = 2;
    private ast.SessionCommands.SessionCloseCommand sessionCloseCommand_;
    /**
     * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
     * @return Whether the sessionCloseCommand field is set.
     */
    @java.lang.Override
    public boolean hasSessionCloseCommand() {
      return ((bitField0_ & 0x00000002) != 0);
    }
    /**
     * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
     * @return The sessionCloseCommand.
     */
    @java.lang.Override
    public ast.SessionCommands.SessionCloseCommand getSessionCloseCommand() {
      return sessionCloseCommand_ == null ? ast.SessionCommands.SessionCloseCommand.getDefaultInstance() : sessionCloseCommand_;
    }
    /**
     * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
     */
    @java.lang.Override
    public ast.SessionCommands.SessionCloseCommandOrBuilder getSessionCloseCommandOrBuilder() {
      return sessionCloseCommand_ == null ? ast.SessionCommands.SessionCloseCommand.getDefaultInstance() : sessionCloseCommand_;
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
      if (((bitField0_ & 0x00000001) != 0)) {
        output.writeMessage(1, getProgramActivity());
      }
      if (((bitField0_ & 0x00000002) != 0)) {
        output.writeMessage(2, getSessionCloseCommand());
      }
      getUnknownFields().writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (((bitField0_ & 0x00000001) != 0)) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, getProgramActivity());
      }
      if (((bitField0_ & 0x00000002) != 0)) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(2, getSessionCloseCommand());
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
      if (!(obj instanceof ast.ProgramOuterClass.Program)) {
        return super.equals(obj);
      }
      ast.ProgramOuterClass.Program other = (ast.ProgramOuterClass.Program) obj;

      if (hasProgramActivity() != other.hasProgramActivity()) return false;
      if (hasProgramActivity()) {
        if (!getProgramActivity()
            .equals(other.getProgramActivity())) return false;
      }
      if (hasSessionCloseCommand() != other.hasSessionCloseCommand()) return false;
      if (hasSessionCloseCommand()) {
        if (!getSessionCloseCommand()
            .equals(other.getSessionCloseCommand())) return false;
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
      if (hasProgramActivity()) {
        hash = (37 * hash) + PROGRAM_ACTIVITY_FIELD_NUMBER;
        hash = (53 * hash) + getProgramActivity().hashCode();
      }
      if (hasSessionCloseCommand()) {
        hash = (37 * hash) + SESSION_CLOSE_COMMAND_FIELD_NUMBER;
        hash = (53 * hash) + getSessionCloseCommand().hashCode();
      }
      hash = (29 * hash) + getUnknownFields().hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static ast.ProgramOuterClass.Program parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static ast.ProgramOuterClass.Program parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static ast.ProgramOuterClass.Program parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static ast.ProgramOuterClass.Program parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static ast.ProgramOuterClass.Program parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static ast.ProgramOuterClass.Program parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static ast.ProgramOuterClass.Program parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static ast.ProgramOuterClass.Program parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static ast.ProgramOuterClass.Program parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static ast.ProgramOuterClass.Program parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static ast.ProgramOuterClass.Program parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static ast.ProgramOuterClass.Program parseFrom(
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
    public static Builder newBuilder(ast.ProgramOuterClass.Program prototype) {
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
     * Protobuf type {@code ast.Program}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessage.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:ast.Program)
        ast.ProgramOuterClass.ProgramOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return ast.ProgramOuterClass.internal_static_ast_Program_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return ast.ProgramOuterClass.internal_static_ast_Program_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                ast.ProgramOuterClass.Program.class, ast.ProgramOuterClass.Program.Builder.class);
      }

      // Construct using ast.ProgramOuterClass.Program.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessage.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessage
                .alwaysUseFieldBuilders) {
          getProgramActivityFieldBuilder();
          getSessionCloseCommandFieldBuilder();
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        bitField0_ = 0;
        programActivity_ = null;
        if (programActivityBuilder_ != null) {
          programActivityBuilder_.dispose();
          programActivityBuilder_ = null;
        }
        sessionCloseCommand_ = null;
        if (sessionCloseCommandBuilder_ != null) {
          sessionCloseCommandBuilder_.dispose();
          sessionCloseCommandBuilder_ = null;
        }
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return ast.ProgramOuterClass.internal_static_ast_Program_descriptor;
      }

      @java.lang.Override
      public ast.ProgramOuterClass.Program getDefaultInstanceForType() {
        return ast.ProgramOuterClass.Program.getDefaultInstance();
      }

      @java.lang.Override
      public ast.ProgramOuterClass.Program build() {
        ast.ProgramOuterClass.Program result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public ast.ProgramOuterClass.Program buildPartial() {
        ast.ProgramOuterClass.Program result = new ast.ProgramOuterClass.Program(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        onBuilt();
        return result;
      }

      private void buildPartial0(ast.ProgramOuterClass.Program result) {
        int from_bitField0_ = bitField0_;
        int to_bitField0_ = 0;
        if (((from_bitField0_ & 0x00000001) != 0)) {
          result.programActivity_ = programActivityBuilder_ == null
              ? programActivity_
              : programActivityBuilder_.build();
          to_bitField0_ |= 0x00000001;
        }
        if (((from_bitField0_ & 0x00000002) != 0)) {
          result.sessionCloseCommand_ = sessionCloseCommandBuilder_ == null
              ? sessionCloseCommand_
              : sessionCloseCommandBuilder_.build();
          to_bitField0_ |= 0x00000002;
        }
        result.bitField0_ |= to_bitField0_;
      }

      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof ast.ProgramOuterClass.Program) {
          return mergeFrom((ast.ProgramOuterClass.Program)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(ast.ProgramOuterClass.Program other) {
        if (other == ast.ProgramOuterClass.Program.getDefaultInstance()) return this;
        if (other.hasProgramActivity()) {
          mergeProgramActivity(other.getProgramActivity());
        }
        if (other.hasSessionCloseCommand()) {
          mergeSessionCloseCommand(other.getSessionCloseCommand());
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
                    getProgramActivityFieldBuilder().getBuilder(),
                    extensionRegistry);
                bitField0_ |= 0x00000001;
                break;
              } // case 10
              case 18: {
                input.readMessage(
                    getSessionCloseCommandFieldBuilder().getBuilder(),
                    extensionRegistry);
                bitField0_ |= 0x00000002;
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
      private int bitField0_;

      private ast.ProgramActivityOuterClass.ProgramActivity programActivity_;
      private com.google.protobuf.SingleFieldBuilder<
          ast.ProgramActivityOuterClass.ProgramActivity, ast.ProgramActivityOuterClass.ProgramActivity.Builder, ast.ProgramActivityOuterClass.ProgramActivityOrBuilder> programActivityBuilder_;
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       * @return Whether the programActivity field is set.
       */
      public boolean hasProgramActivity() {
        return ((bitField0_ & 0x00000001) != 0);
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       * @return The programActivity.
       */
      public ast.ProgramActivityOuterClass.ProgramActivity getProgramActivity() {
        if (programActivityBuilder_ == null) {
          return programActivity_ == null ? ast.ProgramActivityOuterClass.ProgramActivity.getDefaultInstance() : programActivity_;
        } else {
          return programActivityBuilder_.getMessage();
        }
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       */
      public Builder setProgramActivity(ast.ProgramActivityOuterClass.ProgramActivity value) {
        if (programActivityBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          programActivity_ = value;
        } else {
          programActivityBuilder_.setMessage(value);
        }
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       */
      public Builder setProgramActivity(
          ast.ProgramActivityOuterClass.ProgramActivity.Builder builderForValue) {
        if (programActivityBuilder_ == null) {
          programActivity_ = builderForValue.build();
        } else {
          programActivityBuilder_.setMessage(builderForValue.build());
        }
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       */
      public Builder mergeProgramActivity(ast.ProgramActivityOuterClass.ProgramActivity value) {
        if (programActivityBuilder_ == null) {
          if (((bitField0_ & 0x00000001) != 0) &&
            programActivity_ != null &&
            programActivity_ != ast.ProgramActivityOuterClass.ProgramActivity.getDefaultInstance()) {
            getProgramActivityBuilder().mergeFrom(value);
          } else {
            programActivity_ = value;
          }
        } else {
          programActivityBuilder_.mergeFrom(value);
        }
        if (programActivity_ != null) {
          bitField0_ |= 0x00000001;
          onChanged();
        }
        return this;
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       */
      public Builder clearProgramActivity() {
        bitField0_ = (bitField0_ & ~0x00000001);
        programActivity_ = null;
        if (programActivityBuilder_ != null) {
          programActivityBuilder_.dispose();
          programActivityBuilder_ = null;
        }
        onChanged();
        return this;
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       */
      public ast.ProgramActivityOuterClass.ProgramActivity.Builder getProgramActivityBuilder() {
        bitField0_ |= 0x00000001;
        onChanged();
        return getProgramActivityFieldBuilder().getBuilder();
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       */
      public ast.ProgramActivityOuterClass.ProgramActivityOrBuilder getProgramActivityOrBuilder() {
        if (programActivityBuilder_ != null) {
          return programActivityBuilder_.getMessageOrBuilder();
        } else {
          return programActivity_ == null ?
              ast.ProgramActivityOuterClass.ProgramActivity.getDefaultInstance() : programActivity_;
        }
      }
      /**
       * <code>.ast.ProgramActivity program_activity = 1;</code>
       */
      private com.google.protobuf.SingleFieldBuilder<
          ast.ProgramActivityOuterClass.ProgramActivity, ast.ProgramActivityOuterClass.ProgramActivity.Builder, ast.ProgramActivityOuterClass.ProgramActivityOrBuilder> 
          getProgramActivityFieldBuilder() {
        if (programActivityBuilder_ == null) {
          programActivityBuilder_ = new com.google.protobuf.SingleFieldBuilder<
              ast.ProgramActivityOuterClass.ProgramActivity, ast.ProgramActivityOuterClass.ProgramActivity.Builder, ast.ProgramActivityOuterClass.ProgramActivityOrBuilder>(
                  getProgramActivity(),
                  getParentForChildren(),
                  isClean());
          programActivity_ = null;
        }
        return programActivityBuilder_;
      }

      private ast.SessionCommands.SessionCloseCommand sessionCloseCommand_;
      private com.google.protobuf.SingleFieldBuilder<
          ast.SessionCommands.SessionCloseCommand, ast.SessionCommands.SessionCloseCommand.Builder, ast.SessionCommands.SessionCloseCommandOrBuilder> sessionCloseCommandBuilder_;
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       * @return Whether the sessionCloseCommand field is set.
       */
      public boolean hasSessionCloseCommand() {
        return ((bitField0_ & 0x00000002) != 0);
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       * @return The sessionCloseCommand.
       */
      public ast.SessionCommands.SessionCloseCommand getSessionCloseCommand() {
        if (sessionCloseCommandBuilder_ == null) {
          return sessionCloseCommand_ == null ? ast.SessionCommands.SessionCloseCommand.getDefaultInstance() : sessionCloseCommand_;
        } else {
          return sessionCloseCommandBuilder_.getMessage();
        }
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       */
      public Builder setSessionCloseCommand(ast.SessionCommands.SessionCloseCommand value) {
        if (sessionCloseCommandBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          sessionCloseCommand_ = value;
        } else {
          sessionCloseCommandBuilder_.setMessage(value);
        }
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       */
      public Builder setSessionCloseCommand(
          ast.SessionCommands.SessionCloseCommand.Builder builderForValue) {
        if (sessionCloseCommandBuilder_ == null) {
          sessionCloseCommand_ = builderForValue.build();
        } else {
          sessionCloseCommandBuilder_.setMessage(builderForValue.build());
        }
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       */
      public Builder mergeSessionCloseCommand(ast.SessionCommands.SessionCloseCommand value) {
        if (sessionCloseCommandBuilder_ == null) {
          if (((bitField0_ & 0x00000002) != 0) &&
            sessionCloseCommand_ != null &&
            sessionCloseCommand_ != ast.SessionCommands.SessionCloseCommand.getDefaultInstance()) {
            getSessionCloseCommandBuilder().mergeFrom(value);
          } else {
            sessionCloseCommand_ = value;
          }
        } else {
          sessionCloseCommandBuilder_.mergeFrom(value);
        }
        if (sessionCloseCommand_ != null) {
          bitField0_ |= 0x00000002;
          onChanged();
        }
        return this;
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       */
      public Builder clearSessionCloseCommand() {
        bitField0_ = (bitField0_ & ~0x00000002);
        sessionCloseCommand_ = null;
        if (sessionCloseCommandBuilder_ != null) {
          sessionCloseCommandBuilder_.dispose();
          sessionCloseCommandBuilder_ = null;
        }
        onChanged();
        return this;
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       */
      public ast.SessionCommands.SessionCloseCommand.Builder getSessionCloseCommandBuilder() {
        bitField0_ |= 0x00000002;
        onChanged();
        return getSessionCloseCommandFieldBuilder().getBuilder();
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       */
      public ast.SessionCommands.SessionCloseCommandOrBuilder getSessionCloseCommandOrBuilder() {
        if (sessionCloseCommandBuilder_ != null) {
          return sessionCloseCommandBuilder_.getMessageOrBuilder();
        } else {
          return sessionCloseCommand_ == null ?
              ast.SessionCommands.SessionCloseCommand.getDefaultInstance() : sessionCloseCommand_;
        }
      }
      /**
       * <code>optional .ast.SessionCloseCommand session_close_command = 2;</code>
       */
      private com.google.protobuf.SingleFieldBuilder<
          ast.SessionCommands.SessionCloseCommand, ast.SessionCommands.SessionCloseCommand.Builder, ast.SessionCommands.SessionCloseCommandOrBuilder> 
          getSessionCloseCommandFieldBuilder() {
        if (sessionCloseCommandBuilder_ == null) {
          sessionCloseCommandBuilder_ = new com.google.protobuf.SingleFieldBuilder<
              ast.SessionCommands.SessionCloseCommand, ast.SessionCommands.SessionCloseCommand.Builder, ast.SessionCommands.SessionCloseCommandOrBuilder>(
                  getSessionCloseCommand(),
                  getParentForChildren(),
                  isClean());
          sessionCloseCommand_ = null;
        }
        return sessionCloseCommandBuilder_;
      }

      // @@protoc_insertion_point(builder_scope:ast.Program)
    }

    // @@protoc_insertion_point(class_scope:ast.Program)
    private static final ast.ProgramOuterClass.Program DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new ast.ProgramOuterClass.Program();
    }

    public static ast.ProgramOuterClass.Program getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Program>
        PARSER = new com.google.protobuf.AbstractParser<Program>() {
      @java.lang.Override
      public Program parsePartialFrom(
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

    public static com.google.protobuf.Parser<Program> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Program> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public ast.ProgramOuterClass.Program getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_ast_Program_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_ast_Program_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\rprogram.proto\022\003ast\032\026program_activity.p" +
      "roto\032\026session_commands.proto\"\221\001\n\007Program" +
      "\022.\n\020program_activity\030\001 \001(\0132\024.ast.Program" +
      "Activity\022<\n\025session_close_command\030\002 \001(\0132" +
      "\030.ast.SessionCloseCommandH\000\210\001\001B\030\n\026_sessi" +
      "on_close_commandB\006Z\004/astb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          ast.ProgramActivityOuterClass.getDescriptor(),
          ast.SessionCommands.getDescriptor(),
        });
    internal_static_ast_Program_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_ast_Program_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_ast_Program_descriptor,
        new java.lang.String[] { "ProgramActivity", "SessionCloseCommand", });
    descriptor.resolveAllFeaturesImmutable();
    ast.ProgramActivityOuterClass.getDescriptor();
    ast.SessionCommands.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
